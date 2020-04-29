package main

import (
	"fmt"
	"strings"
)

func main() {
	//Input string contains the prefix
	res := strings.HasPrefix("abcdef", "ab")
	fmt.Println(res)

	//Input string doesn't contain the prefix
	res = strings.HasPrefix("abcdef", "ac")
	fmt.Println(res)
}
