package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"
	"os"
	"path/filepath"

	"github.com/zserge/webview"
)

// TODO: remove need for global variables
// rename fileCache?
var fileState map[string]file

var ViewFolder = "views"
var RuntimeJsPath = "runtime"
var JsFolder = "js"
var CssFolder = "styles"

// need not be global now
var root string

// foreign function interface data
// from js to send messages via json
var ffiData map[string]interface{}

// data is stringified json from js
// type is necessary in that json, other keys are different for each type
func handleRPC(w webview.WebView, data string) {
	// this can crash if not sent a json stringified format
	// does not crash for wrong json only non json stringified data
	err := json.Unmarshal([]byte(data), &ffiData)

	checkErr(err)
	fnType := ffiData["type"]

	fmt.Printf("Action type : %s\n", fnType)

	// TODO: standardize all these actions
	// and format
	switch {
	case fnType == "alert":
		w.Dialog(webview.DialogTypeAlert, 0, "title", ffiData["msg"].(string))
	case fnType == "onload":
	case fnType == "load_js":
		fileName := ffiData["fileName"].(string)
		w.Dispatch(func() {
			err := w.Eval(string(fileState[fileName].data))
			checkErr(err)
		})
	case fnType == "load_css":
		fileName := ffiData["fileName"].(string)
		cssData := string(fileState[fileName].data)
		w.Dispatch(func() {
			// Inject CSS
			w.Eval(fmt.Sprintf(`(function(css){
				var style = document.createElement('style');
				var head = document.head || document.getElementsByTagName('head')[0];
				style.setAttribute('type', 'text/css');
				if (style.styleSheet) {
					style.styleSheet.cssText = css;
				} else {
					style.appendChild(document.createTextNode(css));
				}
				head.appendChild(style);
				})("%s")`, template.JSEscapeString(cssData)))
		})
	case fnType == "open_file":
		filePath := w.Dialog(webview.DialogTypeOpen, 0, "Open file", "")
		_ = filePath
	case fnType == "open_dir":
		directoryPath := w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", "")
		_ = directoryPath
	case fnType == "save_file":
		savePath := w.Dialog(webview.DialogTypeSave, 0, "Save file", "")
		_ = savePath
	// default is not being reached
	default:
		fmt.Errorf("No such action %s", fnType)
	}
}

func addFileToState(path string, info os.FileInfo, err error) error {
	_, fileName := filepath.Split(path)

	if !isBlackListed(fileName) {

		data := getFileData(path)

		// remove null check and make them globally
		// for now
		if fileState == nil {
			fileState = make(map[string]file)
		}

		fileState[fileName] = file{fileName, path, data}
	}

	return nil
}

func main() {

	root := getRootPath()

	filepath.Walk(filepath.Join(root, ViewFolder), addFileToState)
	filepath.Walk(filepath.Join(root, RuntimeJsPath), addFileToState)
	filepath.Walk(filepath.Join(root, JsFolder), addFileToState)
	filepath.Walk(filepath.Join(root, CssFolder), addFileToState)

	if index_view, ok := fileState["index.html"]; ok {
		w := webview.New(webview.Settings{
			URL: `data:text/html,` + url.PathEscape(string(index_view.data)),
			ExternalInvokeCallback: handleRPC,
			Debug:     true,
			Resizable: true,
		})

		defer w.Exit()

		w.Run()
	} else {
		panic("index.html has to be present in views folder")
	}

}
