package main

import (
	"fmt"
	"strconv"
	"strings"
)

var pl = fmt.Println

func multiple(x int) (int, bool, string) {
	a := 3 * 3
	b := x > 3
	c := strings.Repeat(strconv.Itoa(x), 3)
	return a, b, c
}

func many(nn ...int) string {
	ss := make([]string, 0)
	for _, v := range nn {
		ss = append(ss, strconv.Itoa(v))
	}
	return strings.Join(ss, "-")
}

func changeValue(x *int) int {
	*x += 10
	return *x
}

func sliceAppender(ss ...string) []string {
	ss = append(ss, "koira")
	ss = append(ss, "apina")
	return ss
}

func main() {
	// pl(multiple(2))
	// pl(multiple(22))
	pl(many(1, 2, 3, 4, 5))
	var s []string
	s = append(s, "kissa")
	pl(s)

	a := 100
	pl(a)
	changeValue(&a)
	pl(a)
	pl(sliceAppender(s...))
}
