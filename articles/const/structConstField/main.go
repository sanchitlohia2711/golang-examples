package main

import "fmt"

type employee struct {
	name string
	age  int
}

func main() {
	const e = employee{
		name: "John",
		age:  21,
	}
	fmt.Println(e)
}
