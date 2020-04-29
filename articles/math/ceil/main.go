package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Ceil(1.6)
	fmt.Println(res)

	res = math.Ceil(-1.6)
	fmt.Println(res)

	res = math.Ceil(1)
	fmt.Println(res)
}
