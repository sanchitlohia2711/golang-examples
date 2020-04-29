package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Copy a file
func main() {
	// Open original file
	original, err := os.Open("original.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()

	// Create new file
	new, err := os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer new.Close()

	//This will copy
	bytesWritten, err := io.Copy(new, original)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytes Written: %d\n", bytesWritten)
}
