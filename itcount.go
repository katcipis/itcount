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
		val := count()
		counterVal.Set("innerText", val)
		time.Sleep(500 * time.Millisecond)
	}
}

func count() string {
	const DAY_DURATION = 24 * time.Hour

	year := 2019
	month := time.January
	day := 30
	hour := 12
	min := 0
	sec := 0
	nsec := 0
	location := time.Local

	deadline := time.Date(year, month, day, hour, min, sec, nsec, location)
	// WHY: Timezone stuff dont work well on js syscall, so I cant load proper location =(
	// The solution for now is to compensate my timezone manually =(
	const tzdiff = time.Hour * 2
	timeUntilGoing := time.Until(deadline) + tzdiff

	days := int64(math.Floor(timeUntilGoing.Hours() / 24.0))
	timeUntilGoing -= time.Duration(time.Duration(days) * DAY_DURATION)

	hours := int64(math.Floor(timeUntilGoing.Hours()))
	timeUntilGoing -= time.Hour * time.Duration(hours)

	minutes := int64(math.Floor(timeUntilGoing.Minutes()))
	timeUntilGoing -= time.Minute * time.Duration(minutes)

	seconds := int64(math.Floor(timeUntilGoing.Seconds()))

	return fmt.Sprintf(
		`
giorni  : %d
ore     : %d
minuti  : %d
secondi : %d`,
		days,
		hours,
		minutes,
		seconds)
}
