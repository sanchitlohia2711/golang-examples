package main

import (
	"fmt"
	"strings"
)

func main() {
	res := strings.ToLower("ABC")
	fmt.Println(res)

	res = strings.ToLower("ABC12$a")
	fmt.Println(res)
}
