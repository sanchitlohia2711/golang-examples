package main

import (
	"fmt"
	"strings"
)

func main() {
	//Output will be 1 as "bc" is present in "abcdef" at index 1
	res := strings.Index("abcdef", "bc")
	fmt.Println(res)

	//Output will be 2 as "cd" is present in "abcdefabcdef" at index 2
	res = strings.Index("abcdefabcdef", "cd")
	fmt.Println(res)

	//Output will be -1 as "ba" is not present in "abcdef"
	res = strings.Index("abcdef", "ba")
	fmt.Println(res)
}
