package main

import "fmt"

var pl = fmt.Println

type (
	TeeSpoon   float64
	TableSpoon float64
	Millilitre float64
)

func tbpToMl(tsp TableSpoon) Millilitre {
	return Millilitre(tsp * 4.92)
}

func main() {
	pl(tbpToMl(10))
}
