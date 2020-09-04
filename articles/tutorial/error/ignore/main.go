package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("non-existing.txt")

	fmt.Println(file)

}
