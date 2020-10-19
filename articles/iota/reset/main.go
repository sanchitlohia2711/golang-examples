package main

import "fmt"

const (
	a = iota
	b
)

const (
	c = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
