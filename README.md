## Knack
Knack is a framework to build cross platform applications using the web stack - 
HTML, CSS and Javascript.
It is built using https://github.com/zserge/webview as the webview.
The size is much smaller compared to electron.

It is under development.

`go get github.com/promignis/knack`

Current usage(**it will change**)

      html/svg  files in views  folder (will pick index.html by default)
      css       files in styles folder
      js        files in js     folder
      image     files in images folder

### Inject
```js
window.onRuntimeLoad = function() {
      _runtime.loadView("contacts.html") // replaces the current html with ./views/contacts.html

      _runtime.loadCss("filename.css") // injects from ./styles/filename.css

      _runtime.loadJs("bundle.js") // injects from ./js/bundle.js

      _runtime.loadImage("test.png", "test") // injects ./images/test.png to img tag with id "test"
}
```

### Files
Open file and get data in callback(will open native file browser)
```js
_runtime.openFile((fileData) => {
      // do something with file data
})
```
Save file with `fileData`

```js
_runtime.saveFile(fileData)
```

### Fuzzy match
Do fuzzy match over a any dictionary
```js
// dict, word, levenshtein_distance, callback

_runtime.fuzzyMatch(["asd", "abc"], "ab", 1, (results) => {
  // returns array of results
  // in this case "abc" as 1 distance away from "ab"
})
```

### Examples
Fuzzy search on files

```js
_runtime.getFileWalker("../", (fileList) => {
  let fileDict = JSON.parse(fileList).map(file => file.name)
  _runtime.fuzzyMatch(fileDict, "main.g", 3, (fuzzyResults) => {
    
  })
})
```

## OSX

### Build
`./build-app.sh`

### Run
`./run-app.sh`
