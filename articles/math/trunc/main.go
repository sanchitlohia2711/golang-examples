package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Trunc(1.6)
	fmt.Println(res)

	res = math.Trunc(-1.6)
	fmt.Println(res)

	res = math.Trunc(1)
	fmt.Println(res)
}
