package main

import "fmt"

var pl = fmt.Println

func main() {
	var heroes map[string]string
	heroes = make(map[string]string)
	heroes["a"] = "sergei"
	heroes["b"] = "stina"
	pl(heroes)
	pl(heroes["a"])

	x, isthere := heroes["x"]
	if isthere {
		pl(x)
	} else {
		pl("no x heroes")
	}

	delete(heroes, "a")
	pl(heroes)

}
