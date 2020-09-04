package main

import "fmt"

func main() {
	switch char := "a"; char {
	case "a":
		fmt.Println("a")
	case "b", "c":
		fmt.Println("b or c")
	default:
		fmt.Println("No matching character")
	}

	//fmt.Println(char)
}
