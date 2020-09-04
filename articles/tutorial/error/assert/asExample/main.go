package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var pathError *os.PathError
	err := openFile("non-existing.txt")

	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("Using Assert: Error e is of type path error. Error: %v\n", e)
	} else {
		fmt.Println("Using Assert: Error not of type path error")
	}

	if errors.As(err, &pathError) {
		fmt.Printf("Using As function: Error e is of type path error. Error: %v\n", pathError)
	}
}

func openFile(fileName string) error {
	_, err := os.Open("non-existing.txt")
	if err != nil {
		return fmt.Errorf("Error opening: %w", err)
	}
	return nil
}
