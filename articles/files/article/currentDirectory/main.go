package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current Wroking Direcoty: %s", currentWorkingDirectory)
}
