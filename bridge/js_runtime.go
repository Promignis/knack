package bridge

import (
	"fmt"

	"html/template"

	"github.com/zserge/webview"
)

// send a value to a js callback
func ResolveJsCallback(w webview.WebView, callbackId int, data string) {
	js := fmt.Sprintf(`_runtime.resolveCallback(%d, "%s")`, callbackId, template.JSEscapeString(data))
	fmt.Println(js)
	RunJsInWebview(w, js)
}

func RunJsInWebview(w webview.WebView, js string) {
	w.Dispatch(func() {
		w.Eval(js)
	})
}

func CssInsertViaJs(cssData string) string {
	return fmt.Sprintf(`(function(css){
				var style = document.createElement('style');
				var head = document.head || document.getElementsByTagName('head')[0];
				style.setAttribute('type', 'text/css');
				if (style.styleSheet) {
					style.styleSheet.cssText = css;
				} else {
					style.appendChild(document.createTextNode(css));
				}
				head.appendChild(style);
				})("%s")`, template.JSEscapeString(cssData))
}
