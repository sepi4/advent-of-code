package main

import (
	"fmt"
	"reflect"
	"time"
)

var pl = fmt.Println

func main() {
	nowtime := time.Now()
	pl(reflect.TypeOf(nowtime))

	a := nowtime.Year()
	pl(a)
	printTime(nowtime)

	lastWeek := nowtime.AddDate(0, 1, -7*20)
	printTime(lastWeek)

}

func printTime(t time.Time) {

	fmt.Printf(
		"%d-%d-%d %d:%d:%d\n",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)

}
