//+build js,wasm
package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("webasm moderfocker")

	document := js.Global().Get("document")
	div := document.Call("getElementById", "target")
	node := document.Call("createElement", "div")
	div.Call("appendChild", node)
	node.Set("innerText", "Hello World")
}
