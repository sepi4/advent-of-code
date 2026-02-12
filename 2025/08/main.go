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
	y       float64
	x       float64
	z       float64
	circuit int
}

type dist struct {
	value float64
	a     *box
	b     *box
}

func getDistance(a box, b box) float64 {
	dx := (a.x - b.x) * (a.x - b.x)
	dy := (a.y - b.y) * (a.y - b.y)
	dz := (a.z - b.z) * (a.z - b.z)
	return math.Sqrt(dx + dy + dz)
}

func changeCircuitId(boxes []box, oldId int, newId int) {
	for i := range boxes {
		if boxes[i].circuit == oldId {
			boxes[i].circuit = newId
		}
	}
}

func allInSameCircuit(boxes []box) bool {
	var pre int = -100
	for i := range boxes {
		if pre == -100 {
			pre = boxes[i].circuit
		} else if boxes[i].circuit != pre {
			return false
		}
	}
	return true
}

func main() {
	// amount := 10
	text := lib.ReadInput("example.txt")

	// amount := 1000
	// text := lib.ReadInput("input.txt")

	rows := lib.GetRows(text)
	var boxes []box
	var distances []dist

	// for _, v := range arr --- SYNTAKSI LUO KOPION ELEMENTISTÄ, JOTEN KÄYTÄ VAIN KUN LUET
	// for i := range arr --- KÄYTÄ INDEKSOINTIA KUN HALUAT MUTATOIDA ELEMENTTEJÄ
	for _, r := range rows {
		rr := strings.Split(r, ",")
		var coor []float64
		for _, s := range rr {
			i, _ := strconv.ParseFloat(s, 64)
			coor = append(coor, i)
		}
		boxes = append(boxes, box{
			x:       coor[0],
			y:       coor[1],
			z:       coor[2],
			circuit: -1,
		})
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

	// sort distances, first is shotest
	slices.SortFunc(distances, func(a, b dist) int {
		return cmp.Compare(a.value, b.value)
	})

	circuitId := 0
	for i := 0; i < len(distances); i++ {
		d := distances[i]
		if d.a.circuit == -1 && d.b.circuit == -1 { // molemmat eivat ole piirissa
			circuitId++
			d.a.circuit = circuitId
			d.b.circuit = circuitId
		} else if d.a.circuit == -1 && d.b.circuit != -1 { // b ei ole piirissa
			d.a.circuit = d.b.circuit
		} else if d.a.circuit != -1 && d.b.circuit == -1 { // a ei ole  piirissa
			d.b.circuit = d.a.circuit
		} else if d.a.circuit != -1 && d.b.circuit != -1 { // molemmat ovat piirissa
			oldId := d.b.circuit
			d.b.circuit = d.a.circuit
			changeCircuitId(boxes, oldId, d.a.circuit)
		}

		if allInSameCircuit(boxes) {
			fmt.Println("part2 ===========")
			fmt.Println("result: ", d.a.x*d.b.x)
			break
		}

	}
}
