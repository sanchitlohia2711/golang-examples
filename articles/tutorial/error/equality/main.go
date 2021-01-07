package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("non-existing.txt")

	if err == os.ErrNotExist {
		fmt.Println("Equality Operator: Both errors are equal")
	} else {
		fmt.Println("Not equal")
	}

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Is function: Both errors are equal")
	}
}
