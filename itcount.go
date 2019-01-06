//+build js,wasm
package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	fmt.Println("webasm moderfocker")

	document := js.Global().Get("document")
	counterVal := document.Call("createElement", "div")

	counterDiv := document.Call("getElementById", "counter")
	counterDiv.Call("appendChild", counterVal)

	for {
		fmt.Println("start loop")
		val := count()
		fmt.Println(val)
		counterVal.Set("innerText", val)
		time.Sleep(time.Second)
	}
}

func count() string {
	year := 2019
	month := time.January
	day := 30
	hour := 0
	min := 0
	sec := 0
	nsec := 0
	location := time.UTC

	deadline := time.Date(year, month, day, hour, min, sec, nsec, location)
	until := time.Until(deadline)
	hours := int(until.Hours())
	minutes := int(until.Minutes())
	seconds := int(until.Seconds())
	days := hours / 24

	return fmt.Sprintf("dias: %d\nhoras: %d\nminutos: %d\nsegundos: %d", days, hours, minutes, seconds)
}
