package main

import (
	"fmt"
	"strings"
)

func main() {

	present := strings.Contains("abc", "ab")
	fmt.Println(present)

	present = strings.Contains("abc", "xyz")
	fmt.Println(present)
}
