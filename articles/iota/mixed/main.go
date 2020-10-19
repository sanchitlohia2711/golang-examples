package main

import "fmt"

const (
	a = iota

	b
	//comment
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
