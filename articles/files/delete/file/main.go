package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
}
