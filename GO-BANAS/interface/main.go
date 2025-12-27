package main

import "fmt"

var pl = fmt.Println

type Animal interface {
	AngrySound()
	HappySound()
	SexSound()
}

type Cat string

func (c Cat) Attack() {
	pl("cat attack my ass")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hisss")
}
func (c Cat) HappySound() {
	pl("Cat says Purrrr")
}
func (c Cat) SexSound() {
	pl("Cat says AAAAAAAAAAAAAA")
}

func main() {
	var kitty Animal
	kitty = Cat("Mirri")
	kitty.AngrySound()
	kitty.SexSound()

}
