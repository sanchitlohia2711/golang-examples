package main

import "fmt"

func main() {
	var v interface{}
	v = "test"

	val, ok := v.(string)
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("v is not of type string")
	}

	val2, ok := v.(int)
	if ok {
		fmt.Println(val2)
	} else {
		fmt.Println("v is not of type string")
	}

}
