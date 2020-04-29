package main

import (
	"fmt"
	"strings"
)

func main() {

	res := strings.Compare("abc", "abc")
	fmt.Println(res)

	res = strings.Compare("abc", "xyz")
	fmt.Println(res)

	res = strings.Compare("xyz", "abc")
	fmt.Println(res)
}
