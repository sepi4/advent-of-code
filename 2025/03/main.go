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
}
func part2() {
}

func main() {
	// part1()
	// part2()
}
