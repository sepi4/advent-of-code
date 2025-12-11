package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

var pl = fmt.Println

func main() {
	pl(rand.Int())
	pl(rand.IntN(10))
	pl(math.Pi)
	pl(math.E)

	i := 10
	for i > 0 {
		pl(rand.IntN(10))
		i--
	}
}
