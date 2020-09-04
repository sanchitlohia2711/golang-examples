package main

import "fmt"

func main() {
	switch age := 45; {
	case age < 18:
		fmt.Println("Kid")
	case age >= 18 && age < 40:
		fmt.Println("Young")
	default:
		fmt.Println("Old")
	}
}
