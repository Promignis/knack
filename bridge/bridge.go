package bridge

import (
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/promignis/knack/fs"
	"github.com/promignis/knack/utils"

	"github.com/zserge/webview"
)

// foreign function interface data
// from js to send messages via json
var ffiData map[string]interface{}

// data is stringified json from js
// type is necessary in that json, other keys are different for each type
func HandleRPC(w webview.WebView, data string) {
	// this can crash if not sent a json stringified format
	// does not crash for wrong json only non json stringified data
	err := json.Unmarshal([]byte(data), &ffiData)

	utils.CheckErr(err)
	fnType := ffiData["type"]

	fmt.Printf("Action type : %s\n", fnType)

	// TODO: standardize all these actions
	// and format
	switch fnType {
	case "alert":
		w.Dialog(webview.DialogTypeAlert, 0, "title", ffiData["msg"].(string))
	case "onload":
	case "load_js":
		fileName := ffiData["fileName"].(string)
		w.Dispatch(func() {
			err := w.Eval(string(fs.FileState[fileName].Data()))
			utils.CheckErr(err)
		})
	case "load_css":
		fileName := ffiData["fileName"].(string)
		cssData := string(fs.FileState[fileName].Data())
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
	case "open_file":
		filePath := w.Dialog(webview.DialogTypeOpen, 0, "Open file", "")
		_ = filePath
	case "open_dir":
		directoryPath := w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", "")
		_ = directoryPath
	case "save_file":
		savePath := w.Dialog(webview.DialogTypeSave, 0, "Save file", "")
		_ = savePath
	// default is not being reached
	default:
		fmt.Errorf("No such action %s", fnType)
	}
}
