package main

import (
	"fmt"
)

//Declare a struct
type employee struct {
	name   string
	age    int
	salary float64
}

func main() {
	//Initialize a struct without named fields
	employee1 := employee{"John", 21, 1000}
	fmt.Println(employee1)

	//Initialize a struct with named fields
	employee2 := employee{
		name:   "Sam",
		age:    22,
		salary: 1100,
	}
	fmt.Println(employee2)

	//Initializing only some fields. Other values are initialized to default zero value of that type
	employee3 := employee{name: "Tina", age: 24}
	fmt.Println(employee3)

}
