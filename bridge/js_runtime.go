package bridge

import (
	"encoding/json"
	"fmt"

	"html/template"

	"github.com/promignis/knack/utils"
	"github.com/zserge/webview"
)

// send a value to a js callback
func ResolveJsCallback(w webview.WebView, cbData *CallbackData) {
	callbackStrData, err := json.Marshal(cbData)
	utils.CheckErr(err)
	js := fmt.Sprintf(`_runtime.resolveCallback("%s")`, template.JSEscapeString(string(callbackStrData)))
	RunJsInWebview(w, js)
}

func RunJsInWebview(w webview.WebView, js string) {
	w.Dispatch(func() {
		w.Eval(js)
	})
}

func InjectCss(cssData string) string {
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

func InjectHtml(htmlData string) string {
	return fmt.Sprintf(`document.open("text/html");
				document.write(`+"`%s`"+`);
				document.close();`, htmlData)
}

func InjectImage(imageData string, imageId string) string {
	return fmt.Sprintf(`document.getElementById("%s").src ="data:image/png;base64,%s";`, imageId, imageData)
}
