if(window){
  (function(){
    window.__runtime = Object.assign(window.__runtime || {}, {
      updateRuntime: function(views, viewData) {
        var views = views ? views : [];
        var viewData = viewData ? viewData : {};
        window.__runtime = Object.assign(window.__runtime, {
          views: views,
          viewData: viewData
        })
      },
      loadJs: function(jsFileName) {
        window.external.invoke(JSON.stringify({type: 'load_js', fileName: jsFileName}))
      },
      alert: function(msg) {
        window.external.invoke(JSON.stringify({type: 'alert', msg: msg}))
      }
    });

    if(window.onRuntimeLoad) {
      window.onRuntimeLoad();
    }

  })();
} else {
  console.error("window not found");
}

