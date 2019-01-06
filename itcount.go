//+build js,wasm
package main

import (
	"fmt"
	"math"
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
		time.Sleep(500 * time.Millisecond)
	}
}

func count() string {
	const DAY_DURATION = 24 * time.Hour

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
	var elapsed time.Duration

	days := int64(math.Floor(until.Hours() / 24.0))
	elapsed += time.Duration(time.Duration(days) * DAY_DURATION)

	hours := int64(math.Floor(time.Duration(until - elapsed).Hours()))
	elapsed += time.Hour * time.Duration(hours)

	minutes := int64(math.Floor(time.Duration(until - elapsed).Minutes()))
	elapsed += time.Minute * time.Duration(minutes)

	seconds := int64(math.Floor(time.Duration(until - elapsed).Seconds()))

	return fmt.Sprintf(
		"giorni: %d\nore: %d\nminuti: %d\nsecondi: %d",
		days,
		hours,
		minutes,
		seconds)
}
