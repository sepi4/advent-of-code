package main

import (
	"fmt"
	"strconv"
	"strings"

	"2025/lib"
)

func part1() {
	// s := lib.ReadInput("06/example.txt")
	s := lib.ReadInput("06/input.txt")
	rows := strings.Split(s, "\n")
	arr := [][]string{}
	for i, r := range rows {
		if i == len(rows)-1 {
			break
		}
		vv := strings.Fields(r)
		arr = append(arr, vv)
	}
	numbers := [][]int64{}
	operations := arr[len(arr)-1]
	for i := 0; i < len(arr)-1; i++ {
		nn := []int64{}
		for j := 0; j < len(arr[i]); j++ {
			x, _ := strconv.ParseInt(arr[i][j], 10, 64)
			nn = append(nn, x)
		}
		numbers = append(numbers, nn)
	}

	results := []int64{}
	for i := range operations {
		op := operations[i]
		num := numbers[0][i]
		for j := 1; j < len(numbers); j++ {
			switch op {
			case "*":
				num *= numbers[j][i]
			case "+":
				num += numbers[j][i]
			default:
				panic("VIRHE")
			}
		}
		results = append(results, num)
	}

	// fmt.Println("numbers:", numbers)
	// fmt.Println("operations:", operations)
	// fmt.Println("results:", results)
	sum := int64(0)
	for _, x := range results {
		sum += x
	}
	fmt.Println("sum:", sum)
}

func part2() {
	// s := lib.ReadInput("06/example.txt")
	s := lib.ReadInput("06/input.txt")
	rows := lib.GetRows(s)
	// fmt.Println("rows:", rows)
	// for _, r := range rows {
	// 	fmt.Println(r)
	// }

	arr := rows[:len(rows)-1]
	var numbers [][]rune
	for _, r := range arr {
		numbers = append(numbers, []rune(r))
	}
	ops := []rune(rows[len(rows)-1])

	allNums := []int64{}
	currentNums := []int64{}
	for x := len(ops) - 1; x >= 0; x-- {
		// loop nums
		num := ""
		for y := 0; y < len(numbers); y++ {
			v := numbers[y][x]
			if v != ' ' {
				if len(num) > 0 {
					num += string(v)
				} else {
					num = string(v)
				}
			}
		}

		// add num to currentNums
		// a, _ := strconv.Atoi(num)
		a, _ := strconv.ParseInt(num, 10, 64)
		currentNums = append(currentNums, a)
		num = ""

		// check ops
		if ops[x] != ' ' {
			currentResult := currentNums[0]
			for i := 1; i < len(currentNums); i++ {
				if ops[x] == '*' {
					currentResult *= currentNums[i]
				} else {
					currentResult += currentNums[i]
				}
			}
			allNums = append(allNums, currentResult)
			x-- // jump over one column
			// reset
			currentNums = []int64{}
		}
	}

	sum := int64(0)
	for _, v := range allNums {
		sum += v
	}
	fmt.Println("sum:", sum)
}

func main() {
	// part1()
	part2()
}
