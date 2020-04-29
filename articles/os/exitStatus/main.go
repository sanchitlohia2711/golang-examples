package main

import (
	"fmt"
	"os"
)

func main() {
	success := true
	if success {
		fmt.Println("Success")
		os.Exit(0)
	} else {
		fmt.Println("Failure")
		os.Exit(1)
	}
}
