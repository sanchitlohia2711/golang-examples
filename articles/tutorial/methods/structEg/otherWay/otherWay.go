package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func (e employee) details() {
	fmt.Printf("Name: %s\n", e.name)
	fmt.Printf("Age: %d\n", e.age)
}

func (e *employee) setName(newName string) {
	e.name = newName
}

func main() {
	emp := employee{name: "Sam", age: 31, salary: 2000}
	employee.details(emp)

	(*employee).setName(&emp, "John")

	fmt.Printf("Name: %s\n", emp.name)
}
