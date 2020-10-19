package main

import "fmt"

type Size int

const (
	small = Size(iota)
	medium
	large
	extraLarge
)

func main() {
	var m Size = 1
	m.String()
}

func (s Size) String() {
	switch s {
	case small:
		fmt.Println("Small")
	case medium:
		fmt.Println("Medium")
	case large:
		fmt.Println("Large")
	case extraLarge:
		fmt.Println("Extra Large")
	default:
		fmt.Println("Invalid Size entry")
	}
}
