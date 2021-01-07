package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("non-existing.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name() + "opened succesfully")
	}
}
