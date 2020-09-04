package main

import "fmt"

func main() {
	age := 45
	switch {
	case age < 18:
		fmt.Println("Kid")
	case age >= 18 && age < 40:
		fmt.Println("Young")
	default:
		fmt.Println("Old")
	}
}
