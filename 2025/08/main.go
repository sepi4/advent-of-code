package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"2025/lib"
)

type box struct {
	y float64
	x float64
	z float64
}

type dist struct {
	value float64
	a     *box
	b     *box
}

type circuit struct {
	boxes []*box
}

func getDistance(a box, b box) float64 {
	dx := (a.x - b.x) * (a.x - b.x)
	dy := (a.y - b.y) * (a.y - b.y)
	dz := (a.z - b.z) * (a.z - b.z)
	return math.Sqrt(dx + dy + dz)
}

func isInCircuit(circ circuit, box *box) bool {
	for i := 0; i < len(circ.boxes); i++ {
		if circ.boxes[i] == box {
			return true
		}
	}
	return false
}

func main() {
	// text := lib.ReadInput("example.txt")
	text := lib.ReadInput("input.txt")

	rows := lib.GetRows(text)
	var boxes []box
	var distances []dist

	for _, r := range rows {
		rr := strings.Split(r, ",")
		var coor []float64
		for _, s := range rr {
			i, _ := strconv.ParseFloat(s, 64)
			coor = append(coor, i)
		}
		boxes = append(boxes, box{coor[0], coor[1], coor[2]})
	}

	for i := 0; i < len(boxes)-1; i++ {
		for j := i + 1; j < len(boxes); j++ {
			distances = append(distances, dist{
				getDistance(boxes[i], boxes[j]),
				&boxes[i],
				&boxes[j],
			})
		}
	}

	slices.SortFunc(distances, func(a, b dist) int {
		return -cmp.Compare(a.value, b.value)
	})

	var circuits []circuit
	// for range 10 { // only 10 in example
	for range boxes { // all boxes in input
		last := distances[len(distances)-1]
		distances = distances[:len(distances)-1]
		addNew := true
		// check if a or b is in some circuit
		// - if is and other box to that circuit also
		// - else make new circuit and append to circuits list
		for ii := 0; ii < len(circuits); ii++ {
			if isInCircuit(circuits[ii], last.a) {
				circuits[ii].boxes = append(circuits[ii].boxes, last.b)
				addNew = false
				break
			} else if isInCircuit(circuits[ii], last.b) {
				circuits[ii].boxes = append(circuits[ii].boxes, last.a)
				addNew = false
				break
			}
		}
		if addNew {
			c := circuit{}
			c.boxes = append(c.boxes, last.a)
			c.boxes = append(c.boxes, last.b)
			circuits = append(circuits, c)
		}
	}

	slices.SortFunc(circuits, func(a, b circuit) int {
		return -cmp.Compare(len(a.boxes), len(b.boxes))
	})

	for _, c := range circuits {
		fmt.Println("len: ", len(c.boxes))
	}
	fmt.Println("----")

	result := 1
	for i := range 3 {
		l := len(circuits[i].boxes)
		result *= l
		fmt.Println(i, ":", l)
	}

	// fmt.Println(len(rows))
	// fmt.Println(rows)
	// fmt.Println(boxes)
	// fmt.Println(circuits)

	fmt.Println("dist: ", len(distances))
	fmt.Println(result)
}
