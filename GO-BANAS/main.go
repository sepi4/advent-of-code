package main

import (
	"fmt"
	"log"

	stuff "example/project/mypackage"
)

var pl = fmt.Println

func main() {
	date := stuff.Date{}
	err := date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(2000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1st day: %d.%d.%d\n", date.Day(), date.Month(), date.Year())

}
