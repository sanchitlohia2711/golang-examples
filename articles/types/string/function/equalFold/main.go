package main

import (
	"fmt"
	"strings"
)

func main() {
	res := strings.EqualFold("abc", "ABC")
	fmt.Println(res)

	res = strings.EqualFold("abc", "aBC")
	fmt.Println(res)

	res = strings.EqualFold("abc", "AbC")
	fmt.Println(res)

	res = strings.EqualFold("abc", "AbCd")
	fmt.Println(res)
}
