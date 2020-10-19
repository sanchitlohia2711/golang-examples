package main

import "fmt"

const (
	a = iota
	b = iota + 4
	c = iota * 4
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
