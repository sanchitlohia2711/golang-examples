package main

import "fmt"

func main() {
	b := newBuilder()

	d := newDirector(b)

	h := d.buildHouse()

	fmt.Printf("House Door Type: %s", h.doorType)
	fmt.Println()
	fmt.Printf("House Window Type: %s", h.windowType)
	fmt.Println()
	fmt.Printf("House Num Floor: %d", h.floor)
	fmt.Println()
}
