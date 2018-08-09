package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/zserge/webview"
)

// TODO: remove need for global variables

// "filename.html": view-file
var view_state map[string]view

// "filename.js" : js-file
var js_state map[string]js

var view_folder = "/views"
var runtime_js_path = "/runtime"
var js_folder = "/js"

// need not be global now
var root string

// foreign function interface data
// from js to send messages via json
var ffiData map[string]interface{}

// data is stringified json from js
// type is necessary in that json, other keys are different for each type
func handleRPC(w webview.WebView, data string) {
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
			err := w.Eval(string(js_state[fileName].data))
			checkErr(err)
		})
	case fnType == "open_file":
		file_path := w.Dialog(webview.DialogTypeOpen, 0, "Open file", "")
		_ = file_path
	case fnType == "open_dir":
		directory_path := w.Dialog(webview.DialogTypeOpen, webview.DialogFlagDirectory, "Open directory", "")
		_ = directory_path
	case fnType == "save_file":
		save_path := w.Dialog(webview.DialogTypeSave, 0, "Save file", "")
		_ = save_path
	}
}

// TODO: refactor
func getWalker(walkerType string) func(string, os.FileInfo, error) error {
	switch {
	case walkerType == "view":
		return func(path string, info os.FileInfo, err error) error {
			_, file_name := filepath.Split(path)

			if !isBlackListed(file_name) {

				data := getFileData(path)

				// remove null check and make them globally
				// for now
				if view_state == nil {
					view_state = make(map[string]view)
				}

				view_state[file_name] = view{file_name, path, data}
			}
			return nil
		}
	case walkerType == "js":
		return func(path string, info os.FileInfo, err error) error {
			_, file_name := filepath.Split(path)

			if !isBlackListed(file_name) {

				data := getFileData(path)
				if js_state == nil {
					js_state = make(map[string]js)
				}

				js_state[file_name] = js{file_name, path, data}
			}
			return nil
		}
	}

	return func(string, os.FileInfo, error) error { return nil }
}

func main() {

	root := getRootPath()
	filepath.Walk(filepath.Join(root, view_folder), getWalker("view"))
	filepath.Walk(filepath.Join(root, runtime_js_path), getWalker("js"))
	filepath.Walk(filepath.Join(root, js_folder), getWalker("js"))

	if index_view, ok := view_state["index.html"]; ok {
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
