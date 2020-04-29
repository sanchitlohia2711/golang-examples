package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "true"
	if val, err := strconv.ParseBool(input); err == nil {
		fmt.Printf("%T, %v\n", val, val)
	}

	input = "false"
	if val, err := strconv.ParseBool(input); err == nil {
		fmt.Printf("%T, %v\n", val, val)
	}

	input = "garbage"
	if val, err := strconv.ParseBool(input); err == nil {
		fmt.Printf("%T, %v\n", val, val)
	} else {
		fmt.Println("Given input is not a bool")
	}
}
