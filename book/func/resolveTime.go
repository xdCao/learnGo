package main

import "fmt"

const (
	SecondsPerMinute = 60
	SecondsPerHour   = SecondsPerMinute * 60
	SecondsPerDay    = SecondsPerHour * 24
)

func main() {

	fmt.Println(resolveTime(1000))

	_, hour, minute := resolveTime(18000)

	fmt.Println(hour, minute)

	day, _, _ := resolveTime(90000)
	fmt.Println(day)

}

func resolveTime(seconds int) (day int, hours int, minute int) {
	day = seconds / SecondsPerDay
	hours = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}
