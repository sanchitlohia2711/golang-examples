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

func (e employee) getSalary() int {
	return e.salary
}

func main() {
	emp := employee{name: "Sam", age: 31, salary: 2000}
	emp.details()

	fmt.Printf("Salary %d\n", emp.getSalary())
}
