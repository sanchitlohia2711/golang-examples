package main

import "fmt"

func main() {
	switch char := "b"; {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
		break
		fmt.Println("after b")
	default:
		fmt.Println("No matching character")
	}
}
