package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
