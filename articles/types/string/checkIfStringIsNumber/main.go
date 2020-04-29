package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "1234"
	val, err := strconv.Atoi(x)
	if err != nil {
		fmt.Printf("Supplied value %s is not a number\n", x)
	} else {
		fmt.Printf("Supplied value %s is a number with value %d\n", x, val)
	}

	y := "123b"
	val, err = strconv.Atoi(y)
	if err != nil {
		fmt.Printf("Supplied value %s is not a number\n", y)
	} else {
		fmt.Printf("Supplied value %s is a number with value %d\n", y, val)
	}
}
