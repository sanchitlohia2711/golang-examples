package main

import (
	"fmt"
	"strings"
)

func main() {
	//It will replaces 1 instance of a with 1
	res := strings.Replace("abcdabxyabr", "a", "1", 1)
	fmt.Println(res)

	//It will replace all instances of a with 1
	res = strings.Replace("abcdabxyabr", "a", "", -1)
	fmt.Println(res)
}
