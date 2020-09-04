package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func (e *employee) setNewName(newName string) {
	e.name = newName
}

func main() {
	emp := employee{name: "Sam", age: 31, salary: 2000}
	emp.setNewName("John")

	fmt.Printf("Name: %s\n", emp.name)

	(&emp).setNewName("Mike")

	fmt.Printf("Name: %s\n", emp.name)

}
