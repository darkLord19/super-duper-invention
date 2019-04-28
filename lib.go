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

func getMaleName(this js.Value, args []js.Value) interface{} {
	rand := randomdata.FullName(randomdata.Male)
	return nil
}

func getFemaleName(this js.Value, args []js.Value) interface{} {
	rand := randomdata.FullName(randomdata.Female)
	return nil
}

func getSillyName(this js.Value, args []js.Value) interface{} {
	rand := randomdata.SillyName()
	return nil
}

func getRandomPara(this js.Value, args []js.Value) interface {}{
	rand := randomdata.Paragraph()
	return nil
}

func registerFuncs() {
	setSillyName = js.FuncOf(getSillyName)
	setMaleName = js.FuncOf(getMaleName)
	setFemaleName = js.FuncOf(getFemaleName)
	setRandomPara = js.FuncOf(getRandomPara)
}

func main() {
	c := make(chan struct{}, 0)
	println("WASM Go Initialized")
	registerFuncs()
	<-c
}
