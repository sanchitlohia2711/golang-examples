package main

import "fmt"

type employee struct {
	Name   string `json:"n"`
	Age    int    `json:"a"`
	Salary int    `json:"s"`
}

func main() {

	emp := employee{Name: "Sam", Age: 31, Salary: 2000}

	fmt.Printf("Emp: %v\n", emp)
	fmt.Printf("Emp: %+v\n", emp)
	fmt.Printf("Emp: %#v\n", emp)

	fmt.Println(emp)
}
