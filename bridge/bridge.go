package bridge

import (
	"encoding/json"
	"fmt"

	"github.com/promignis/knack/fs"
	"github.com/promignis/knack/utils"
	"github.com/promignis/knack/validation"

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
		RunJsInWebview(w, string(fs.FileState[fileName].Data()))
	case "load_css":
		fileName := ffiData["fileName"].(string)
		cssData := string(fs.FileState[fileName].Data())
		RunJsInWebview(w, CssInsertViaJs(cssData))
	case "open_file":
		filePath := w.Dialog(webview.DialogTypeOpen, 0, "Open file", "")
		// transfer binary data as natively as possible
		// profile for speed
		data := string(fs.GetFileData(filePath))
		callbackId := int(ffiData["callbackId"].(float64))
		ResolveJsCallback(w, callbackId, data)
	case "open_dir":
		directoryPath := w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", "")
		_ = directoryPath
	case "save_file":
		savePath := w.Dialog(webview.DialogTypeSave, 0, "Save file", "")
		fileData := ffiData["fileData"].(string)
		// check if path is valid
		if validation.IsValidPath(savePath) {
			fs.WriteFileData(savePath, []byte(fileData))
		}
	// default is not being reached
	default:
		fmt.Errorf("No such action %s", fnType)
	}
}
