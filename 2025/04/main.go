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

func collectToBeRemoved(rows []string) [][]int {
	limit := 4
	arr := make([][]int, 0)
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[0]); x++ {
			if rows[y][x] == '@' && getNextAmount(rows, y, x) < limit {
				arr = append(arr, []int{y, x})
			}
		}
	}
	return arr
}

func makeRemove(rows []string, arr [][]int) []string {
	for _, v := range arr {
		y, x := v[0], v[1]
		rr := []rune(rows[y])
		rr[x] = '.'
		rows[y] = string(rr)
	}
	return rows
}

func part2() {
	// s := lib.ReadInput("04/example.txt")
	s := lib.ReadInput("04/input.txt")
	sum := 0
	rows := lib.GetRows(s)
	toRemove := collectToBeRemoved(rows)
	// fmt.Println(rows)
	for len(toRemove) > 0 {
		// make remove
		sum += len(toRemove)
		rows = makeRemove(rows, toRemove)
		toRemove = collectToBeRemoved(rows)
	}
	// fmt.Println(rows)
	fmt.Println(sum)
}

func main() {
	// part1()
	part2()
}
