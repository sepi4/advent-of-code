package main

import (
	"fmt"
	"slices"

	"2025/lib"
)

type Stream struct {
	y          int
	x          int
	splitCount int
}

const SPLITTER = '^'

func checkSamePosition(arr []*Stream, y int, x int) []*Stream {
	aa := []*Stream{}
	for _, v := range arr {
		if v.y == y && v.x == x {
			aa = append(aa, v)
		}
	}
	newaa := []*Stream{}
	for _, v := range arr {
		if !slices.Contains(aa, v) {
			newaa = append(newaa, v)
		}
	}

	if len(aa) == 2 {
		newaa = append(newaa, &Stream{
			y:          aa[0].y,
			x:          aa[0].x,
			splitCount: aa[0].splitCount + aa[1].splitCount,
		})
		return newaa
	}
	return arr
}

func part1() {
	s := lib.ReadInput("07/example.txt")
	// s := lib.ReadInput("07/input.txt")
	rr := lib.GetRows(s)
	rows := [][]rune{}
	for _, r := range rr {
		rows = append(rows, []rune(r))
	}

	// for _, r := range rows {
	// 	fmt.Println(r)
	// }

	// FIND START
	startIndex := -1
	for i, x := range rows[0] {
		if x == 'S' {
			startIndex = i
			break
		}
	}

	ss := []*Stream{{
		y:          1,
		x:          startIndex,
		splitCount: 0,
	}}

	moved := true
	for moved {
		moved = false
		// MOVE STREAMS
		for i := 0; i < len(ss); i++ {
			if rows[ss[i].y][ss[i].x] != SPLITTER {
				rows[ss[i].y][ss[i].x] = '|'
			}
			if ss[i].y == len(rows)-1 {
				continue
			}
			// move down current
			ss[i].y++
			moved = true
			// add another to right
			if rows[ss[i].y][ss[i].x] == SPLITTER {
				s2 := &Stream{
					y:          ss[i].y,
					x:          ss[i].x + 1,
					splitCount: 0,
				}
				ss = append(ss, s2)
				ss = checkSamePosition(ss, ss[i].y, ss[i].x)
				ss[i].x-- // move current to left
				ss[i].splitCount++
				// if len(aaa) > 1 {
				// 	fmt.Println("aaa", aaa)
				// 	return
				// }
			}
			ss = checkSamePosition(ss, ss[i].y, ss[i].x)
		}
	}

	// count splits
	count := 0
	for _, s := range ss {
		count += s.splitCount
	}

	// for _, r := range rows {
	// 	fmt.Println(string(r))
	// }
	fmt.Println("ss:", ss)
	fmt.Println("count:", count)
}

func part2() {
}

func main() {
	part1()
	// part2()
}
