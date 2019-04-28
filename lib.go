// +build js, wasm

package main

import (
	"syscall/js"

	"github.com/Pallinder/go-randomdata"
)

var (
	document      js.Value
	setSillyName  js.Func
	setMaleName   js.Func
	setFemaleName js.Func
	setRandomPara js.Func
	clearAll      js.Func
)

func init() {
	document = js.Global().Get("document")
}

func getElementByID(id string) js.Value {
	return document.Call("getElementById", id)
}

func addEventListener(id string, cb js.Func) {
	getElementByID(id).Call("addEventListener", "click", cb)
}

func addInnerText(text string, id string) {
	getElementByID(id).Set("innerText", text)
}

func addClass(class string, id string) {
	getElementByID(id).Get("classList").Call("add", class)
}

func removeClass(class string, id string) {
	getElementByID(id).Get("classList").Call("remove", class)
}

func getMaleName(this js.Value, args []js.Value) interface{} {
	rand := randomdata.FullName(randomdata.Male)
	addInnerText(rand, "maleT")
	removeClass("scale-out", "maleT")
	addClass("scale-in", "maleT")
	return nil
}

func getFemaleName(this js.Value, args []js.Value) interface{} {
	rand := randomdata.FullName(randomdata.Female)
	addInnerText(rand, "femaleT")
	removeClass("scale-out", "femaleT")
	addClass("scale-in", "femaleT")
	return nil
}

func getSillyName(this js.Value, args []js.Value) interface{} {
	rand := randomdata.SillyName()
	addInnerText(rand, "sillyT")
	removeClass("scale-out", "sillyT")
	addClass("scale-in", "sillyT")
	return nil
}

func getRandomPara(this js.Value, args []js.Value) interface {}{
	rand := randomdata.Paragraph()
	addInnerText(rand, "textT")
	removeClass("scale-out", "textT")
	addClass("scale-in", "textT")
	return nil
}

func clearAllHeadings(this js.Value, args []js.Value) interface {} {
	ids := []string{"textT", "sillyT", "maleT", "femaleT"}
	for i:=0; i<len(ids); i++ {
		removeClass("scale-in", ids[i])
		addClass("scale-out", ids[i])
	}
	return nil
}

func registerFuncs() {
	setSillyName = js.FuncOf(getSillyName)
	setMaleName = js.FuncOf(getMaleName)
	setFemaleName = js.FuncOf(getFemaleName)
	setRandomPara = js.FuncOf(getRandomPara)
	clearAll = js.FuncOf(clearAllHeadings)
}

func addEventListenerHelper() {
	addEventListener("silly", setSillyName)
	addEventListener("male", setMaleName)
	addEventListener("female", setFemaleName)
	addEventListener("text", setRandomPara)
	addEventListener("clear", clearAll)
}

func main() {
	c := make(chan struct{}, 0)
	println("WASM Go Initialized")
	registerFuncs()
	addEventListenerHelper()
	<-c
}
