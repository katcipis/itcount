//+build js,wasm
package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("webasm moderfocker")

	document := js.Global().Get("document")
	counterVal := document.Call("createElement", "div")

	counter := document.Call("getElementById", "counter")
	counter.Call("appendChild", counterVal)
	counterVal.Set("innerText", "Hello World")
}
