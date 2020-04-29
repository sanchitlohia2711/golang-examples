package main

import (
	"fmt"
	"strings"
)

func main() {
	res := strings.ToUpper("abc")
	fmt.Println(res)

	res = strings.ToUpper("abc12$")
	fmt.Println(res)
}
