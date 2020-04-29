package main

import (
	"fmt"
	"strings"
)

func main() {
	//Case 1 Input string contains the substring
	res := strings.Count("abcdabcd", "ab")
	fmt.Println(res)

	//Case 1 Input string doesn't contains the substring
	res = strings.Count("abcdabcd", "xy")
	fmt.Println(res)

	//Case 1 Substring supplied is an empty string
	res = strings.Count("abcdabcd", "")
	fmt.Println(res)
}
