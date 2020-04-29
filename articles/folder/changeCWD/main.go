package main

import (
	"fmt"
	"os"
)

func main() {
	os.Chdir("/Users")

	newDir, err := os.Getwd()
	if err != nil {

	}
	fmt.Printf("Current Working Direcotry:  %s", newDir)
}
