package main

import "fmt"

const (
	a = iota + 10
	b
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
