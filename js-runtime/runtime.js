if(window){
  (function(){
    window._runtime = Object.assign(window._runtime || {}, JsRuntime());

    if(window.onRuntimeLoad) {
      window.onRuntimeLoad();
    }

  })();
} else {
  console.error("window not found");
}

function sendAction(actionData) {
  window.external.invoke(JSON.stringify(actionData))
}

// returns the runtime functions in object that will be merged
// with _runtime
function JsRuntime(){
  return {
    // {cbId: function}
    callbackIds: {},
    // load javascript via fileName
    loadJs: function(jsFileName) {
      sendAction({type: 'load_js', fileName: jsFileName})
    },
    alert: function(msg) {
      sendAction({type: 'alert', msg: msg})
    },
    resolveCallback: function(resolveObjJSONStr) {
      const resolveObj = JSON.parse(resolveObjJSONStr)
      const cbId = resolveObj.CallbackId
      const args = resolveObj.Args
      _runtime.callbackIds[cbId].apply(this, args)
      // can have issues?
      delete _runtime.callbackIds[cbId]
    },
    // if callback is empty returns promise
    openFile: function(cb) {
      if(!cb){
        return new Promise((resolve, reject) => {
          sendAction({type: 'open_file', callbackId: _runtime.getCbId(resolve)})
        })
      }
      sendAction({type: 'open_file', callbackId: _runtime.getCbId(cb)})
    },
    saveFile: function(fileData) {
      sendAction({type: 'save_file', fileData: fileData})
    },
    getCbId: function(cb) {
      var cbLen = Object.keys(_runtime.callbackIds).length
      _runtime.callbackIds[cbLen] = cb
      return cbLen
    },
    loadCss: function(cssFileName) {
      sendAction({type: 'load_css', fileName: cssFileName})
    },
    loadView: function(viewName) {
      sendAction({type: 'load_html', fileName: viewName})
    },
    loadImage: function(imageName, imageId) {
      sendAction({type: 'load_img', imageName, imageId})
    },
    getFileWalker: function(filePath, cb) {
      return new FileWalker(filePath, cb)
    },
    getFileStat: function(filePath, cb) {
      return new FileStat(filePath, cb)
    },
    // if callback is empty then returns promise
    fuzzyMatch: function(dict, word, distance, cb) {
      if(!cb) {
        return new Promise((resolve, reject) => {
          sendAction({type: 'fuzzy_match', dict: JSON.stringify(dict), word, distance, callbackId: _runtime.getCbId(resolve)})
        })
      }
      sendAction({type: 'fuzzy_match', dict: JSON.stringify(dict), word, distance, callbackId: _runtime.getCbId(cb)})
    }
  }
}

// turn to more functional approach later
function FileWalker(filePath, cb) {
  this.filePath = filePath

  this.cbId = _runtime.getCbId(cb)

  this.fileStat = new FileStat(filePath, (fileStat) => {
    this.fileStat = fileStat
  })
  this.walk()
}

FileWalker.prototype.walk = function() {
  sendAction({type: 'file_walker', filePath: this.filePath, callbackId: this.cbId})
}

// get stat for file path
function FileStat(filePath, cb) {
  this.filePath = filePath
  this.cbId = _runtime.getCbId(cb)
  this.getFileStat()
}

FileStat.prototype.getFileStat = function() {
  sendAction({type: 'file_stat', filePath: this.filePath, callbackId: this.cbId}, (fileStat) => {
    this.fileStat = fileStat
  })
}
