package main

import (
	"fmt"
	"strings"
)

func main() {
	res := strings.Replace("abcdabxyabr", "ab", "12", 1)
	fmt.Println(res)

	res = strings.Replace("abcdabxyabr", "ab", "12", -1)
	fmt.Println(res)
}
