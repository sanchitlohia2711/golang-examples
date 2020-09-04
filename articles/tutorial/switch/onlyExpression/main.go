package main

import "fmt"

func main() {
	char := "a"
	switch char {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	default:
		fmt.Println("No matching character")
	}
}
