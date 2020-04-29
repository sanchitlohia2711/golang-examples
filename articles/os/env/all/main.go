package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//Set env a to b
	err := os.Setenv("a", "b")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Setenv("c", "d")
	if err != nil {
		log.Fatal(err)
	}

	//Get all env variables
	fmt.Println(os.Environ())

	//Clear all env variables
	os.Clearenv()

	//Again get all env variable
	fmt.Println(os.Environ())
}
