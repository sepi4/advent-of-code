package main

import (
	"fmt"
	"slices"
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

type FreshAvailable struct {
	fresh     [][]int64
	available []int64
}

func getFreshAndAvailable(rows []string) FreshAvailable {
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
	return FreshAvailable{fresh, available}
}

func part1() {
	// s := lib.ReadInput("05/example.txt")
	s := lib.ReadInput("05/input.txt")
	rows := strings.Split(s, "\n")
	value := getFreshAndAvailable(rows)

	// check
	freshAvailable := []int64{}
	for _, x := range value.available {
		if check(x, value.fresh) {
			freshAvailable = append(freshAvailable, x)
		}
	}
	fmt.Println(len(freshAvailable))
}

func isBetween(x int64, arr []int64) bool {
	return x >= arr[0] && x <= arr[1]
}

func count(arr [][]int64) int64 {
	var sum int64 = 0
	for _, vv := range arr {
		x := vv[1] - vv[0]
		if x < 0 {
			// panic("less than zero!")
			fmt.Println(vv)
			continue
		}
		sum += x + 1
	}
	return sum
}

/*
2-8 6-10 ... 1 false
2-8 6-10 ... 2  true

6-10 2-8 ... 1  true
6-10 2-8 ... 2  false

---

11-12 1-20 ... 1 false
11-12 1-20 ... 2 true

1-20 11-12 ... 1 false
1-20 11-12 ... 2 false
*/
func InEachOther(a []int64, b []int64) bool {
	if a[0] >= b[0] && a[0] <= b[1] {
		return true
	}
	if a[1] >= b[0] && a[1] <= b[1] {
		return true
	}
	if b[0] >= a[0] && b[0] <= a[1] {
		return true
	}
	if b[1] >= a[0] && b[1] <= a[1] {
		return true
	}
	return false
}

func part2() {
	// s := lib.ReadInput("05/example.txt")
	// s := lib.ReadInput("05/example2.txt")
	s := lib.ReadInput("05/input.txt")
	rows := strings.Split(s, "\n")
	value := getFreshAndAvailable(rows)

	fresh := value.fresh
	for i := 0; i < len(fresh); i++ {
		// GET COMBINED
		toCombineIndexes := []int{}
		for j := i + 1; j < len(fresh); j++ {
			if InEachOther(fresh[i], fresh[j]) {
				toCombineIndexes = append(toCombineIndexes, j)
			}
		}
		// RESOLVE COMBINED
		if len(toCombineIndexes) > 0 {
			toCombineIndexes = append(toCombineIndexes, i)
			newFresh := [][]int64{}
			toCombine := []int64{}
			for a := 0; a < len(fresh); a++ {
				if slices.Contains(toCombineIndexes, a) {
					toCombine = append(toCombine, fresh[a][0], fresh[a][1])
				} else {
					newFresh = append(newFresh, fresh[a])
				}
			}

			// ADD MIN MAX
			newFresh = append(newFresh, []int64{
				slices.Min(toCombine),
				slices.Max(toCombine),
			})

			// NEW FRESH, START LOOP AGAIN
			i = -1
			fresh = newFresh
		}

	}

	// GET SUM
	sum := int64(0)
	for _, arr := range fresh {
		v := arr[1] - arr[0]
		sum += v + 1
	}
	fmt.Println(sum)
}

func main() {
	// part1()
	part2()
}
