package main

import (
	"fmt"
	"strings"
)

func main() {
	//Case 1 s contains sep. Will output slice of length 3
	res := strings.Join([]string{"ab", "cd", "ef"}, "-")
	fmt.Println(res)

	//Case 2 slice is empty. It will output a empty string
	res = strings.Join([]string{}, "-")
	fmt.Println(res)

	//Case 3 sep is empty. It will output a string combined from the slice of strings
	res = strings.Join([]string{"ab", "cd", "ef"}, "")
	fmt.Println(res)
}
