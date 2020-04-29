package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//Set env a to b
	err := os.Setenv("a", "x")
	if err != nil {
		log.Fatal(err)
	}

	val, present := os.LookupEnv("a")
	fmt.Printf("a env variable present: %t\n", present)
	fmt.Println(val)

	val, present = os.LookupEnv("b")
	fmt.Printf("b env variable present: %t\n", present)
}
