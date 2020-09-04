package main

import "fmt"

func main() {
	age := 8

	if age < 18 {
		fmt.Println("Kid")
	} else if age >= 18 && age < 40 {
		fmt.Println("Young")
	} else {
		fmt.Println("Old")
	}
}
