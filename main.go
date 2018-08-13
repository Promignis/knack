package main

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/promignis/knack/bridge"
	"github.com/promignis/knack/constants"
	"github.com/promignis/knack/fs"
	"github.com/promignis/knack/utils"

	"github.com/zserge/webview"
)

// TODO: remove need for global variables
// rename fileCache?

// need not be global now
var root string

func addFileToState(path string, info os.FileInfo, err error) error {
	_, fileName := filepath.Split(path)

	if !utils.IsBlackListed(fileName) {

		data := fs.GetFileData(path)

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

	if indexView, ok := fs.FileState[constants.DefaultIndexFile]; ok {
		w := webview.New(webview.Settings{
			URL: `data:text/html,` + url.PathEscape(string(indexView.Data())),
			ExternalInvokeCallback: bridge.HandleRPC,
			Debug:     true,
			Resizable: true,
		})

		if runtimeJs, ok := fs.FileState[constants.RuntimeJsFile]; ok {
			bridge.RunJsInWebview(w, string(runtimeJs.Data()))
		} else {
			panic("failed to load runtime.js")
		}

		defer w.Exit()

		w.Run()
	} else {
		panic("index.html has to be present in views folder")
	}

}
