var input = document.getElementById('alertText')
var btn = document.getElementById('btn')
var changeViewBtn = document.getElementById('changeView')
var currentText = ""

input.addEventListener("change", function(ev) {
  currentText = input.value
}, false)

btn.addEventListener('click', function(e) {
  _runtime.alert(currentText)
}, false)

changeViewBtn.addEventListener('click', function(e) {
  _runtime.loadView('index2.html')
}, false)

