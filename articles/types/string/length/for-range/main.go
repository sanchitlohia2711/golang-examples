package main

import "fmt"

func main() {
	sample := "a£b"

	len := 0
	//With index and value
	fmt.Println("Both Index and Value")
	for i, letter := range sample {
		len++
		fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))
	}

	fmt.Printf("Length of given string is %d\n", len)
}
