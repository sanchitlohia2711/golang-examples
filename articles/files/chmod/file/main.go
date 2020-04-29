package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create new file
	new, err := os.Create("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer new.Close()

	stats, err := os.Stat("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Permission File Before: %s\n", stats.Mode())

	err = os.Chmod("new.txt", 0700)
	if err != nil {
		log.Fatal(err)
	}
	stats, err = os.Stat("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Permission File After: %s\n", stats.Mode())
}
