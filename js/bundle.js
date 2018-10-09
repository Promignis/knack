var input = document.getElementById('alertText')
var btn = document.getElementById('btn')
var changeViewBtn = document.getElementById('changeView')
var persistToFileBtn = document.getElementById('persistToFile')
var getFromFileBtn = document.getElementById('getFromFile')

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

persistToFileBtn.addEventListener('click', function() {
  _runtime.setToFile("testfile", JSON.stringify({a:1}))
})

getFromFileBtn.addEventListener('click', function() {
  _runtime.getFromFile("testfile", function(stringifiedJson) {
    alert(stringifiedJson)
    console.log(JSON.parse(stringifiedJson))
  })  
})

