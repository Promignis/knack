package main

import (
	"net/url"
	"path/filepath"

	"github.com/promignis/knack/bridge"
	"github.com/promignis/knack/constants"
	"github.com/promignis/knack/fs"
	"github.com/promignis/knack/utils"

	"github.com/zserge/webview"
)

func main() {

	root := utils.GetRootPath()

	// scope them in state
	// to prevent name clashes
	filepath.Walk(filepath.Join(root, constants.ViewFolder), fs.AddFileToState)
	filepath.Walk(filepath.Join(root, constants.RuntimeJsPath), fs.AddFileToState)
	filepath.Walk(filepath.Join(root, constants.JsFolder), fs.AddFileToState)
	filepath.Walk(filepath.Join(root, constants.CssFolder), fs.AddFileToState)
	filepath.Walk(filepath.Join(root, constants.ImageFoler), fs.AddFileToState)

	if indexView, ok := fs.FileState[constants.DefaultIndexFile]; ok {
		w := webview.New(webview.Settings{
			URL: `data:text/html,` + url.PathEscape(indexView.StringData()),
			ExternalInvokeCallback: bridge.HandleRPC,
			Debug:     true,
			Resizable: true,
		})

		if runtimeJs, ok := fs.FileState[constants.RuntimeJsFile]; ok {
			bridge.RunJsInWebview(w, runtimeJs.StringData())
		} else {
			panic("Failed to load runtime.js")
		}

		defer w.Exit()

		w.Run()
	} else {
		panic("File index.html has to be present in views folder")
	}

}
