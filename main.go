package main

import (
	"net/url"
	"os"
	"path/filepath"

	"knack/bridge"
	"knack/constants"
	"knack/fs"
	"knack/utils"

	"github.com/zserge/webview"
)

// TODO: remove need for global variables
// rename fileCache?

// need not be global now
var root string

func addFileToState(path string, info os.FileInfo, err error) error {
	_, fileName := filepath.Split(path)

	if !utils.IsBlackListed(fileName) {

		data := utils.GetFileData(path)

		// remove null check and make them globally
		// for now
		if fs.FileState == nil {
			fs.FileState = make(map[string]*fs.File)
		}

		fs.FileState[fileName] = fs.NewFile(fileName, path, data)
	}

	return nil
}

func main() {

	root := utils.GetRootPath()

	filepath.Walk(filepath.Join(root, constants.ViewFolder), addFileToState)
	filepath.Walk(filepath.Join(root, constants.RuntimeJsPath), addFileToState)
	filepath.Walk(filepath.Join(root, constants.JsFolder), addFileToState)
	filepath.Walk(filepath.Join(root, constants.CssFolder), addFileToState)

	if index_view, ok := fs.FileState["index.html"]; ok {
		w := webview.New(webview.Settings{
			URL: `data:text/html,` + url.PathEscape(string(index_view.Data())),
			ExternalInvokeCallback: bridge.HandleRPC,
			Debug:     true,
			Resizable: true,
		})

		defer w.Exit()

		w.Run()
	} else {
		panic("index.html has to be present in views folder")
	}

}
