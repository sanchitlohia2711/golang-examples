package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := writeToTempFile("Some text")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Write to file succesful")
}

func writeToTempFile(text string) error {
	file, err := os.Open("temp.txt")
	if err != nil {
		return err
	}

	n, err := file.WriteString("Some text")
	if err != nil {
		return err
	}

	fmt.Printf("N bytes written: %d", n)

	file.Close()
	return nil
}
