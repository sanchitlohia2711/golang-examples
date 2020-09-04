package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	emp := employee{"Sam", 31, 2000}

	fmt.Printf("Emp: %+v\n", emp)

	//emp = employee{"Sam", 31}
}
