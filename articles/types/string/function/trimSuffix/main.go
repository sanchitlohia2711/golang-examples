package main

import (
	"fmt"
	"strings"
)

func main() {
	//This will output removed
	res := strings.TrimSuffix("removedtest", "test")
	fmt.Println(res)

	//The input string will remain unchanged as it doesn't contain the test as suffix
	res2 := strings.TrimSuffix("removedtes", "test")
	fmt.Println(res2)
}
