## Knack
Knack is a framework to build cross platform applications using the web stack - 
HTML, CSS and Javascript.
It is built using https://github.com/zserge/webview as the webview.
The size is much smaller compared to electron.

It is under development.

`go get github.com/promignis/knack`

Current usage(it will change)
Place html files in `views`  folder (will pick `index.html` default)
      css  files in `styles` folder
      js   files in `js`     folder

### Inject
`_runtime.loadCss("filename.css")`
`_runtime.loadJs("bundle.js")`

### Files
Open file and get data in callback(will open native file browser)
`_runtime.openFile((fileData) => {
      // do something with file data
})`

## OSX

### Build
`./build-app.sh`

### Run
`./run-app.sh`
