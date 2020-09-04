package main

import "fmt"

func main() {
	switch "a" {
	case "a":
		fmt.Println("a")
	case "a":
		fmt.Println("Another a")
	case "b":
		fmt.Println("b")
	default:
		fmt.Println("No matching character")
	}
}
