package main

import (
	"syscall/js"

	"github.com/Pallinder/go-randomdata"
)

var (
	document      js.Value
	setSillyName  js.Callback
	setMaleName   js.Callback
	setFemaleName js.Callback
	setRandomPara js.Callback
)

func init() {
	document = js.Global().Get("document")
}

func setMaleName() {
	rand := randomdata.FullName(randomdata.Male)
}

func setFemaleName() {
	rand := radomdata.FullName(randomdata.Female)
}

func setSillyName() {
	rand := randomdata.SillyName()
}

func setRandomPara() {
	rand := randomdata.Paragraph()
}

func registerCallbacks() {
	setSillyName = js.Global().Set("setSillyName", js.NewCallback(setRandomSillyName))
	setMaleName = js.Global().Set("setMaleName", js.NewCallback(setRandomSillyName))
	setFemaleName = js.Global().Set("setFemaleName", js.NewCallback(setRandomSillyName))
	setRandomPara = js.Global().Set("setRandomPara", js.NewCallback(setRandomSillyName))
}

func main() {
	c := make(chan struct{}, 0)
	println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
