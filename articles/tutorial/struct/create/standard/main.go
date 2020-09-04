package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	emp1 := employee{}
	fmt.Printf("Emp1: %+v\n", emp1)

	emp2 := employee{name: "Sam", age: 31, salary: 2000}
	fmt.Printf("Emp2: %+v\n", emp2)

	emp3 := employee{
		name:   "Sam",
		age:    31,
		salary: 2000,
	}
	fmt.Printf("Emp3: %+v\n", emp3)

	emp4 := employee{
		name: "Sam",
		age:  31,
	}
	fmt.Printf("Emp4: %+v\n", emp4)
}
