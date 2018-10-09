// TODO: Add webpack to clean up pollyfilling for windows
// Pollyfilling  Object.assign
if (typeof Object.assign != 'function') {
  Object.defineProperty(Object, "assign", {
    value: function assign(target, varArgs) {
      'use strict';
      if (target == null) {
        throw new TypeError('Cannot convert undefined or null to object');
      }

      var to = Object(target);

      for (var index = 1; index < arguments.length; index++) {
        var nextSource = arguments[index];

        if (nextSource != null) { 
          for (var nextKey in nextSource) {
            if (Object.prototype.hasOwnProperty.call(nextSource, nextKey)) {
              to[nextKey] = nextSource[nextKey];
            }
          }
        }
      }
      return to;
    },
    writable: true,
    configurable: true
  });
}

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
    openFile: function(cb) {
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
      sendAction({type: 'load_img', imageName: imageName, imageId: imageId})
    },
    getFileWalker: function(filePath, cb) {
      return new FileWalker(filePath, cb)
      // sendAction({type: 'file_walker', callbackId: _runtime.getCbId(cb), filePath})
    },
    getFileStat: function(filePath, cb) {
      return new FileStat(filePath, cb)
    },
    fuzzyMatch: function(dict, word, distance, cb) {
      sendAction({type: 'fuzzy_match', dict: JSON.stringify(dict), word: word, distance: distance, callbackId: _runtime.getCbId(cb)})
    },
    setToFile: function(filename, stringifiedJson) {
      sendAction({type: 'set_to_file', filename: filename, stringifiedJson: stringifiedJson})
    },
    getFromFile: function(filename, cb) {
      sendAction({type: 'get_from_file', filename: filename, callbackId: _runtime.getCbId(cb)})
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
