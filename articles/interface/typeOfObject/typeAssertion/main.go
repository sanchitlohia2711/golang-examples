package main

import "fmt"

func main() {

	var test interface{}
	test = "test_string"

	val, ok := test.(string)
	if ok {
		fmt.Printf("Test is of type string with value %s\n", val)
	} else {
		fmt.Printf("Unknown Type %T", test)
	}

	test = 2
	val2, ok := test.(int)
	if ok {
		fmt.Printf("Test is of type int with value %d\n", val2)
	} else {
		fmt.Printf("Unknown Type %T", test)
	}
}
