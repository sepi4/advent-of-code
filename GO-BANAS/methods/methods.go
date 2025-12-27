package main

import (
	"fmt"
	"math"
)

var pl = fmt.Println

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func main() {
	r := rect{2, 3}
	pl("Rect area:", r.area())
	c := circle{1.5}
	pl("Circle area:", c.area())
	if true {
		pl("not nil")
	}
}
