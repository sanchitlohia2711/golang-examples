package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type dog struct {
	age int
}

func (d dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog walk")
}

type pet1 struct {
	a    animal
	name string
}

type pet2 struct {
	animal
	name string
}

func main() {
	d := dog{age: 5}

	p1 := pet1{name: "Milo", animal: d}

	fmt.Println(p1.name)
	// p1.breathe()
	// p1.walk()

	p1.a.breathe()
	p1.a.walk()

	p2 := pet2{name: "Milo", animal: d}

	fmt.Println(p1.name)
	p2.breathe()
	p2.walk()

	p2.animal.breathe()
	p2.animal.walk()
}
