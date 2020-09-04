package main

import "fmt"

type employee struct {
	name        string
	age         int
	salary      int
	departments []string
}

func main() {
	emp1 := employee{name: "Sam", age: 31, salary: 2000, departments: []string{"CS"}}

	emp2 := employee{name: "Sam", age: 31, salary: 2000, departments: []string{"EC"}}

	if emp1 == emp2 {
		fmt.Println("emp1 annd emp2 are equal")
	} else {
		fmt.Println("emp1 annd emp2 are not equal")
	}
}
