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

func getMaleName() {
	rand := randomdata.FullName(randomdata.Male)
}

func getFemaleName() {
	rand := radomdata.FullName(randomdata.Female)
}

func getSillyName() {
	rand := randomdata.SillyName()
}

func getRandomPara() {
	rand := randomdata.Paragraph()
}

func registerCallbacks() {
	setSillyName = js.Global().Set("setSillyName", js.NewCallback(getSillyName))
	setMaleName = js.Global().Set("setMaleName", js.NewCallback(getMaleName))
	setFemaleName = js.Global().Set("setFemaleName", js.NewCallback(getFemaleName))
	setRandomPara = js.Global().Set("setRandomPara", js.NewCallback(getRandomPara))
}

func main() {
	c := make(chan struct{}, 0)
	println("WASM Go Initialized")
	registerCallbacks()
	<-c
}
