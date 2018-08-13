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

// returns the runtime functions in object that will be merged
// with _runtime
function JsRuntime() {
  return {
    callbackIds: {},
    updateRuntime: function(views, viewData) {
      var views = views ? views : [];
      var viewData = viewData ? viewData : {};
      window._runtime = Object.assign(window._runtime, {
        views: views,
        viewData: viewData
      })
    },
    loadJs: function(jsFileName) {
      window.external.invoke(JSON.stringify({type: 'load_js', fileName: jsFileName}))
    },
    alert: function(msg) {
      window.external.invoke(JSON.stringify({type: 'alert', msg: msg}))
    },
    resolveCallback: function(callbackId, data) {
      _runtime.callbackIds[callbackId](data)
      delete _runtime.callbackIds[callbackId]
    },
    openFile: function(cb) {
      window.external.invoke(JSON.stringify({type: 'open_file', callbackId: _runtime.getCbId(cb)}))
    },
    saveFile: function(fileData) {
      window.external.invoke(JSON.stringify({type: 'save_file', fileData: fileData}))
    },
    getCbId: function(cb) {
      var cbLen = Object.keys(_runtime.callbackIds).length
      _runtime.callbackIds[cbLen] = cb
      return cbLen
    },
    loadCss: function(cssFileName) {
      window.external.invoke(JSON.stringify({type: 'load_css', fileName: cssFileName}))
    }
  }
}
