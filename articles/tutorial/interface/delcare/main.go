package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

func main() {
	var a animal

	fmt.Println(a)
	fmt.Printf("Underlying Type: %T\n", a)
	fmt.Printf("Underlying Value: %v\n", a)
}
