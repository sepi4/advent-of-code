package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(filePath string) string {
	data, err := os.ReadFile(filePath)
	check(err)
	return string(data)
}

func part1() {
	// input := readInput("./example_02.txt")
	input := readInput("./input_02.txt")
	arr := strings.Split(input, ",")
	pl(arr)

	sum := 0
	for _, v := range arr {
		ss := strings.Split(v, "-")
		a, _ := strconv.Atoi(ss[0])
		b, _ := strconv.Atoi(ss[1])

		for i := a; i < b+1; i++ {
			strI := strconv.Itoa(i)
			middle := len(strI) / 2
			if strI[:middle] == strI[middle:] {
				sum += i
				pl("jooo", i)
			}
		}
	}
	pl(sum)
}
func part2() {
	// input := readInput("./example_02.txt")
	input := readInput("./input_02.txt")
	arr := strings.Split(input, ",")
	pl(arr)

	sum := 0
	for _, v := range arr {
		ss := strings.Split(v, "-")
		a, _ := strconv.Atoi(ss[0])
		b, _ := strconv.Atoi(ss[1])

		for i := a; i < b+1; i++ {
			strI := strconv.Itoa(i)
			if repeating(strI) {
				sum += i
				// pl("jooo", i)
			}
		}
	}
	pl(sum)
}
func repeating(s string) bool {
	for subLen := len(s) / 2; subLen > 0; subLen-- {
		if len(s)%subLen != 0 {
			continue
		}
		pre := s[0:subLen]
		isSame := true
		for j := subLen; j < len(s); j += subLen {
			sub := s[j : j+subLen]
			if pre != sub {
				isSame = false
				break
			}
		}
		if isSame {
			return true
		}
	}
	return false
}

func main() {
	// part1()
	part2()
	// x := repeating("2121212121")
	// pl(x)
}
