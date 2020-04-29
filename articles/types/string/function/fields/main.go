package main

import (
	"fmt"
	"strings"
)

func main() {
	//Case 1 Input string contains single spaces
	res := strings.Fields("ab cd ef")
	fmt.Println(res)

	//Case 2 Input string doesn't contain white spaces
	res = strings.Fields("abcdef")
	fmt.Println(res)

	//Case 3 Input string contains double white spaces and spaces at end too.
	res = strings.Fields("ab  cd   ef ")
	fmt.Println(res)

	//Case 4: Input string contain white spaces only. Will output a empty slice
	res = strings.Fields("  ")
	fmt.Println(res)

	//Case 5: Input string is empty. Will output a empty slice
	res = strings.Fields("")
	fmt.Println(res)
}
