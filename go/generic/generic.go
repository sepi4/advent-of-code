package main

import "fmt"

type MyConstraint interface {
	string | int
}

func print[T MyConstraint](a T, b T) int {
	fmt.Println(a)
	fmt.Println(b)
	return 0
}

func main() {
	print("kissa", "kissa")
	print(123, 321)
}
