package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {

	emp := employee{name: "Sam", age: 31, salary: 2000}

	fmt.Printf("Emp: %v\n", emp)
	fmt.Printf("Emp: %+v\n", emp)
	fmt.Printf("Emp: %#v\n", emp)

	fmt.Println(emp)
}
