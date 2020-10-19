package main

import (
	"fmt"
)

func main() {

	a := []string{"a", "b"}
	checkAndPrint(a, 2)
	fmt.Println("Exiting normally")
}

func checkAndPrint(a []string, index int) {
	if index > (len(a) - 1) {
		errorString := fmt.Sprintf("Out of bounds access for slice. Index passed: %d", index)
		panic(errorString)
	}
	fmt.Println(a[index])
}

func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}
