package main

import (
	"fmt"
	"strings"
)

func main() {
	res := strings.ReplaceAll("abcdabxyabr", "ab", "12")
	fmt.Println(res)

	res = strings.ReplaceAll("abcabcabcdef", "ab", "")
	fmt.Println(res)
}
