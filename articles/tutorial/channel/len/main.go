package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 5
	fmt.Printf("Len: %d\n", len(ch))
	fmt.Printf("Capacity: %d\n", cap(ch))

	ch <- 6
	fmt.Printf("Len: %d\n", len(ch))
	fmt.Printf("Capacity: %d\n", cap(ch))

	ch <- 7
	fmt.Printf("Len: %d\n", len(ch))
	fmt.Printf("Capacity: %d\n", cap(ch))
}
