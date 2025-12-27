package main

import (
	"2025/lib"
	"fmt"
	"strconv"
	"strings"
)

var pl = fmt.Println

func printArr(mm []M) {
	for _, x := range mm {
		fmt.Print(x.value)
	}
	fmt.Print("\n")
	fmt.Print("\n")
}

func part1() {
	s := lib.ReadInput("03/example.txt")
	// s := lib.ReadInput("03/input.txt")
	rows := strings.Split(s, "\n")
	maxes := make([]int, len(rows))
	for i, row := range rows {
		maxes[i] = 0
		for a := 0; a < len(row); a++ {
			for b := a + 1; b < len(row); b++ {
				sa := string(row[a])
				sb := string(row[b])
				value, _ := strconv.Atoi(sa + sb)
				if value > maxes[i] {
					maxes[i] = value
				}
			}
		}
	}
	sum := 0
	for i := 0; i < len(maxes)-1; i++ {
		sum += maxes[i]
	}
	// pl(sum)
}

type M struct {
	index int
	value int
}

func toInt64(mm []M) int64 {
	str := ""
	for _, v := range mm {
		str += strconv.Itoa(v.value)
	}
	value, _ := strconv.ParseInt(str, 10, 64)
	return value
}

func filterRows(rows []string) []string {
	arr := make([]string, 0)
	for _, r := range rows {
		if len(r) > 0 {
			arr = append(arr, r)
		}
	}
	return arr
}

func part2() {
	batCount := 12
	// s := lib.ReadInput("03/example.txt")
	s := lib.ReadInput("03/input.txt")
	rows := strings.Split(s, "\n")
	rows = filterRows(rows)

	var sum int64 = 0

	for _, row := range rows {
		// return

		arr := make([]M, 0)
		for len(arr) < batCount {
			m := M{
				index: -1,
				value: -1,
			}
			i := 0
			if len(arr) > 0 {
				i = arr[len(arr)-1].index + 1
			}
			for ; i < len(row)-(batCount-len(arr))+1; i++ {
				v, _ := strconv.Atoi(string(row[i]))
				if v > m.value {
					m = M{
						index: i,
						value: v,
					}
				}
			}
			arr = append(arr, m)
		}
		// pl(toInt64(arr))
		sum += toInt64(arr)
	}
	pl(sum)
}

func main() {
	// part1()
	part2()
}
