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

	//Get env a
	val := os.Getenv("a")
	fmt.Println(val)

	//Unset a
	err = os.Unsetenv("a")
	if err != nil {
		log.Fatal(err)
	}

	val, present := os.LookupEnv("a")
	fmt.Printf("a env variable present: %t\n", present)
}
