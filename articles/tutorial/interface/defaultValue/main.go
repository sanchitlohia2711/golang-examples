package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

func main() {
	var a animal
	fmt.Println(a)
}
