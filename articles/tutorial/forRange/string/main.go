package main

import "fmt"

func main() {
	sample := "a£b"

	//With index and value
	fmt.Println("Both Index and Value")
	for i, letter := range sample {
		fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))
	}

	//Only value
	fmt.Println("\nOnly value")
	for _, letter := range sample {
		fmt.Printf("Value:%s\n", string(letter))
	}

	//Only index
	fmt.Println("\nOnly Index")
	for i := range sample {
		fmt.Printf("Start Index: %d\n", i)
	}
}
