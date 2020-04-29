package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("sample")
	if err != nil {
		log.Fatal(err)
	}
}
