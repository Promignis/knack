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
      sendAction({type: 'load_img', imageName, imageId})
    }
  }
}
