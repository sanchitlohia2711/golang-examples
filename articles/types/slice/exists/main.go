package main

import "fmt"

func main() {
	source := []string{"san", "man", "tan"}

	result := find(source, "san")
	fmt.Printf("Item san found: %t\n", result)

	result = find(source, "can")
	fmt.Printf("Item can found: %t\n", result)
}

func find(source []string, value string) bool {
	for _, item := range source {
		if item == value {
			return true
		}
	}
	return false
}
