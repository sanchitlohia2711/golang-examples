package main

import (
	"fmt"
	"strings"
)

func main() {
	res := strings.NewReplacer("abc", "ABC")
	fmt.Println(res)

}
