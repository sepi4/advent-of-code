package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	var a [5]int
	for _, v := range a {
		pl(v)
	}

	pl("----")
	b := []int{1, 2, 3, 4}
	for _, v := range b {
		pl(v)
	}

	pl("----")
	c := [][]int{
		{1, 2},
		{3, 4, 123},
		{5, 6},
	}
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			fmt.Printf("%d ", c[i][j])
		}
		pl("")
	}

	d := []rune("kissa")
	for _, v := range d {
		fmt.Printf("%c", v)
	}
}
