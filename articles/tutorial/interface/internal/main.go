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

	a = lion{age: 10}
	fmt.Printf("Underlying Type: %T\n", a)
	fmt.Printf("Underlying Value: %v\n", a)
}
