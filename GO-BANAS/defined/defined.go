package main

import "fmt"

var pl = fmt.Println

type (
	TeeSpoon   float64
	TableSpoon float64
	Millilitre float64
)

func teeSpnToMl(teespoon TeeSpoon) Millilitre {
	return Millilitre(teespoon * 4.92)
}

func tableSpnToMl(tablespoon TableSpoon) Millilitre {
	return Millilitre(tablespoon * 14.79)
}

// associative methods
func (teespoon TeeSpoon) ToMls() Millilitre {
	return Millilitre(teespoon * 4.92)
}
func (tablespoon TableSpoon) ToMls() Millilitre {
	return Millilitre(tablespoon * 14.79)
}

func main() {
	a := Millilitre(TeeSpoon(3) * 4.92)
	fmt.Printf("3 tee spoon = %.2f\n", a)

	b := Millilitre(TableSpoon(3) * 14.79)
	fmt.Printf("3 tablespoon = %.2f\n", b)

	c := teeSpnToMl(3)
	fmt.Printf("3 tee spoon = %.2f\n", c)

	pl("2 teespoon + 4 teespoon = ", TeeSpoon(2) > TeeSpoon(4))

	pl("2 teespoon = ", TeeSpoon(2).ToMls())
	pl("2 tablespoon = ", TableSpoon(2).ToMls())
}
