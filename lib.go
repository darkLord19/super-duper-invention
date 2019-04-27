package main

import (
	"syscall/js"

	"github.com/Pallinder/go-randomdata"
)

var document js.Value

func init() {
	document = js.Global().Get("document")
}

func setRandomSillyName() {
	rand := randomdata.SillyName()
	div := document.Call("getElementById", "target")
	node := document.Call("createElement", "div")
	node.Set("innerText", rand)
	div.Call("appendChild", node)
}

func main() {

}
