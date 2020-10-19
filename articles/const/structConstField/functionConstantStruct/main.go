package main

import "fmt"

type employee struct {
	name string
	age  int
}

func main() {
	e := baseEmployee()
	fmt.Println(e)
}

func baseEmployee() employee {
	return employee{
		name: "Unnamed",
		age:  0,
	}
}
