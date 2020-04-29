package main

import (
	"fmt"
	"strings"
)

func main() {
	//Case 1 s contains sep. Will output slice of length 3
	res := strings.Split("ab$cd$ef", "$")
	fmt.Println(res)

	//Case 2 s doesn't contain sep. Will output slice of length 1
	res = strings.Split("ab$cd$ef", "-")
	fmt.Println(res)

	//Case 3 sep is empty. Will output slice of length 8
	res = strings.Split("ab$cd$ef", "")
	fmt.Println(res)

	//Case 4 both s and sep are empty. Will output empty slice
	res = strings.Split("", "")
	fmt.Println(res)

}
