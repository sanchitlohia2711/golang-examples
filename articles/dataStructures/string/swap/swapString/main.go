package main

import "fmt"

func main() {
	a := "123"
	b := "xyz"

	fmt.Printf("Before a:%s b:%s\n", a, b)
	a, b = b, a

	fmt.Printf("After a:%s b:%s\n", a, b)
}
