package main

import (
	"fmt"
	"strings"
)

func main() {
	//Output will be 1 as "bc" is present in "abcdef" at index 1
	res := strings.LastIndex("abcdef", "bc")
	fmt.Println(res)

	//Output will be 8 as "cd" is present in "abcdefabcdef" at index 8
	res = strings.LastIndex("abcdefabcdef", "cd")
	fmt.Println(res)

	//Output will be -1 as "ba" is not present in "abcdef"
	res = strings.LastIndex("abcdef", "ba")
	fmt.Println(res)
}
