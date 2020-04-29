package main

import (
	"fmt"
	"strconv"
)

func main() {
	e1 := "1.3434"
	if s, err := strconv.ParseFloat(e1, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	if s, err := strconv.ParseFloat(e1, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}
