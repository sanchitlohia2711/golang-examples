package main

import (
	"fmt"
)

func main() {

	s := make(chan int, 3)
	fmt.Println(s)
	var test interface{}
	test = "test_string"

	//Using Sprintf
	testType := fmt.Sprintf("%T", test)
	fmt.Println(testType)

	//Using printf
	fmt.Printf("%T\n", test)
}
