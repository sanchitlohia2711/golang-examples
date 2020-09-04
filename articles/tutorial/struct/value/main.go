package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	emp1 := employee{name: "Sam", age: 31, salary: 2000}
	fmt.Printf("Emp1 Before: %v\n", emp1)
	emp2 := emp1
	emp2.name = "John"
	fmt.Printf("Emp1 After assignment: %v\n", emp1)
	fmt.Printf("Emp2: %v\n", emp2)
	test(emp1)
	fmt.Printf("Emp1 After Test Function Call: %v\n", emp1)
}
func test(emp employee) {
	emp.name = "Mike"
	fmt.Printf("Emp in Test function: %v\n", emp)
}
