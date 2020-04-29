package main

import (
	"fmt"
	"strings"
)

func main() {
	//This will output removed
	res := strings.TrimPrefix("testremoved", "test")
	fmt.Println(res)

	//The input string will remain unchanged as it doesn't contain the test as prefix
	res2 := strings.TrimPrefix("tesremoved", "test")
	fmt.Println(res2)
}
