package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

func main() {
	var a animal
	var b animal
	var c animal
	var d animal
	var e animal

	a = lion{age: 10}
	b = lion{age: 10}
	c = lion{age: 5}

	if a == b {
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("a and b are not equal")
	}

	if a == c {
		fmt.Println("a and c are equal")
	} else {
		fmt.Println("a and c are not equal")
	}

	if d == e {
		fmt.Println("d and e are equal")
	} else {
		fmt.Println("d and e are not equal")
	}
}
