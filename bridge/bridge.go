package bridge

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

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
	fnType := string(ffiData["type"].(string))

	fmt.Printf("Action type : %s\n", fnType)

	// TODO: standardize all these actions
	// and format
	// find better way to add than switch case
	switch fnType {
	case "alert":
		w.Dialog(webview.DialogTypeAlert, 0, "title", ffiData["msg"].(string))
	case "load_html":
		fileName := ffiData["fileName"].(string)
		htmlData := fs.FileState[fileName].StringData()
		RunJsInWebview(w, InjectHtml(htmlData))
	case "load_js":
		fileName := ffiData["fileName"].(string)
		RunJsInWebview(w, fs.FileState[fileName].StringData())
	case "load_css":
		fileName := ffiData["fileName"].(string)
		cssData := string(fs.FileState[fileName].StringData())
		RunJsInWebview(w, InjectCss(cssData))
	case "open_file":
		filePath := w.Dialog(webview.DialogTypeOpen, 0, "Open file", "")
		// transfer binary data as natively as possible
		// profile for speed
		data := string(fs.GetFileData(filePath))
		HandleCallback(w, ffiData, []string{data})
	case "open_dir":
		directoryPath := w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", "")
		_ = directoryPath
	case "save_file":
		// savePath should be correct as it is coming from the GUI
		savePath := w.Dialog(webview.DialogTypeSave, 0, "Save file", "")
		fileData := ffiData["fileData"].(string)
		fs.WriteFileData(savePath, []byte(fileData))
	case "load_img":
		imageName := ffiData["imageName"].(string)
		imageId := ffiData["imageId"].(string)
		imageData := fs.FileState[imageName].Data()
		base64Img := base64.StdEncoding.EncodeToString(imageData)
		RunJsInWebview(w, InjectImage(base64Img, imageId))
	case "file_walker":
		filePath := ffiData["filePath"].(string)
		fileList := fs.GetFileList(filePath)
		stringified, err := json.Marshal(fileList)
		utils.CheckErr(err)

		fileData := []string{string(stringified)}

		HandleCallback(w, ffiData, fileData)
	case "file_stat":
		filePath := ffiData["filePath"].(string)
		stat := fs.GetFileStat(filePath)
		// centralize json marshalling and instead of crash due to wrong data being sent
		// better errors
		stringifiedStat, err := json.Marshal(stat)
		utils.CheckErr(err)
		args := []string{string(stringifiedStat)}
		HandleCallback(w, ffiData, args)
	default:
		fmt.Printf("No such action %s", fnType)
	}
}
