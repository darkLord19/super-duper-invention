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

func getMaleName(this js.Value, args []js.Value) interface{} {
	rand := randomdata.FullName(randomdata.Male)
	addInnerText(rand, "maleT")
	return nil
}

func getFemaleName(this js.Value, args []js.Value) interface{} {
	rand := randomdata.FullName(randomdata.Female)
	addInnerText(rand, "femaleT")
	return nil
}

func getSillyName(this js.Value, args []js.Value) interface{} {
	rand := randomdata.SillyName()
	addInnerText(rand, "sillyT")
	return nil
}

func getRandomPara(this js.Value, args []js.Value) interface {}{
	rand := randomdata.Paragraph()
	addInnerText(rand, "textT")
	return nil
}

func registerFuncs() {
	setSillyName = js.FuncOf(getSillyName)
	setMaleName = js.FuncOf(getMaleName)
	setFemaleName = js.FuncOf(getFemaleName)
	setRandomPara = js.FuncOf(getRandomPara)
}

func addEventListenerHelper() {
	addEventListener("silly", setSillyName)
	addEventListener("male", setMaleName)
	addEventListener("female", setFemaleName)
	addEventListener("text", setRandomPara)
}

func main() {
	c := make(chan struct{}, 0)
	println("WASM Go Initialized")
	registerFuncs()
	addEventListenerHelper()
	<-c
}
