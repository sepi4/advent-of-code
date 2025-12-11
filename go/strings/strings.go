package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	pl(reflect.TypeOf(pl))
	pl(reflect.TypeOf(123))

	cv1 := 1.5
	cv2 := int(cv1)
	pl(cv2)

	cv3 := "50000000"
	// convertion
	cv4, err := strconv.Atoi(cv3)

	pl(cv4, err)

	s := "kissa,koira,apina,sika"
	aa := strings.Split(s, ",")
	for v := range aa {
		pl(v)
	}
	b := strings.ReplaceAll(s, ",", "|")
	pl(b)

	c := strings.Contains(s, "koira")
	pl(c)
	d := strings.Contains(s, "kana")
	pl(d)

	pl("-----")
	for i, v := range "kissa istuu" {
		// pl(i, v)
		fmt.Printf("%d %d %c\n", i, v, v)
	}

}
