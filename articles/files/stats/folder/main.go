package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//Create a folder
	err := os.Mkdir("temp", 0755)
	if err != nil {
		log.Fatal(err)
	}

	//Gets stats of the file
	stats, err := os.Stat("temp")
	if err != nil {
		log.Fatal(err)
	}

	//Prints stats of the file
	fmt.Printf("Permission: %s\n", stats.Mode())
	fmt.Printf("Name: %s\n", stats.Name())
	fmt.Printf("Size: %d\n", stats.Size())
	fmt.Printf("Modification Time: %s\n", stats.ModTime())
}
