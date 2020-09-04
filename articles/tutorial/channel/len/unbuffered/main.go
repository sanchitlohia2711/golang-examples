package main

import "fmt"

func main() {
	ch := make(chan int)

	fmt.Printf("Len: %d\n", len(ch))
	fmt.Printf("Capacity: %d\n", cap(ch))
}
