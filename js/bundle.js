var input = document.getElementById('alertText')
var btn = document.getElementById('btn')
var currentText = ""

input.addEventListener("change", function(ev) {
  currentText = input.value
}, false)

btn.addEventListener('click', function(e) {
  window.__runtime.alert(currentText)
}, false)

