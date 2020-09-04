package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {

	emp := employee{name: "Sam", age: 31, salary: 2000}
	empP := &emp
	fmt.Printf("Emp: %+v\n", empP)

	empP = &employee{name: "John", age: 30, salary: 3000}
	fmt.Printf("Emp: %+v\n", empP)
}
