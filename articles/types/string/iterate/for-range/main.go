package main

import "fmt"

func main() {
	sample := "a£b"
	for i, letter := range sample {
		fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))
	}
}
