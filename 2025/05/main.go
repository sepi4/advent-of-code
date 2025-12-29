package main

import (
	"fmt"
	"strconv"
	"strings"

	"2025/lib"
)

func check(x int64, fresh [][]int64) bool {
	for _, v := range fresh {
		if x >= v[0] && x <= v[1] {
			return true
		}
	}
	return false
}

func part1() {
	// s := lib.ReadInput("05/example.txt")
	s := lib.ReadInput("05/input.txt")
	rows := strings.Split(s, "\n")
	fresh := [][]int64{}
	available := []int64{}
	i := 0
	for ; rows[i] != ""; i++ {
		r := rows[i]
		nums := strings.Split(r, "-")
		a, _ := strconv.ParseInt(nums[0], 10, 64)
		b, _ := strconv.ParseInt(nums[1], 10, 64)
		fresh = append(fresh, []int64{a, b})
	}
	for i = i + 1; i+1 < len(rows); i++ {
		r := rows[i]
		x, _ := strconv.ParseInt(r, 10, 64)
		available = append(available, x)
	}
	// fmt.Println(fresh)
	// fmt.Println(available)

	// check
	freshAvailable := []int64{}
	for _, x := range available {
		if check(x, fresh) {
			freshAvailable = append(freshAvailable, x)
		}
	}
	fmt.Println(len(freshAvailable))
}

func main() {
	part1()
	// part2()
}
