package main

import (
	"fmt"
	"strings"
)

func main() {
	//This will output removed
	res := strings.TrimSpace(" test  ")
	fmt.Println(res)

	res = strings.TrimSpace(" This is test  ")
	fmt.Println(res)

}
