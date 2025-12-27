package main

import (
	"2025/lib"
	"fmt"
)

func getNextAmount(rows []string, y int, x int) int {
	arr := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},

		{0, -1},
		{0, 1},

		{1, -1},
		{1, 0},
		{1, 1},
	}

	amount := 0
	for _, v := range arr {
		dy := v[0]
		dx := v[1]
		Y := y + dy
		X := x + dx
		if Y < 0 || X < 0 || Y >= len(rows) || X >= len(rows[0]) {
			continue
		}
		if rows[Y][X] == '@' {
			amount++
		}
	}
	return amount
}

func part1() {
	// s := lib.ReadInput("04/example.txt")
	s := lib.ReadInput("04/input.txt")
	rows := lib.GetRows(s)
	// fmt.Println(rows)
	// fmt.Println(getNextAmount(rows, 0, 1))

	limit := 4
	sum := 0
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[0]); x++ {
			if rows[y][x] == '@' && getNextAmount(rows, y, x) < limit {
				sum++
			}
		}
	}
	fmt.Println("sum:", sum)
}

func main() {
	part1()
}
