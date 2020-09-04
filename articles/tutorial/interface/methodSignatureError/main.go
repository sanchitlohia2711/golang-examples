package main

import "fmt"

type animal interface {
	breathe()
	walk()
	speed() int
}

type lion struct {
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

func (l lion) speed() {
	fmt.Println("Lion walk")
}

func main() {
	var a animal

	a = lion{}
	a.breathe()
	a.walk()
}
