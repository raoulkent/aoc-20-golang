package main

import (
	"fmt"
	"os"

	days "./days"
)

func main() {
	whichDay(os.Args[1])
}

func whichDay(day string) {
	switch day {
	case "1":
		fmt.Println("Running day 1")
		days.One()
	default:
		fmt.Println("Pick a day between 1 and 25")
	}
}
