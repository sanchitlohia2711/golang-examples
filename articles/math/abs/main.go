package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Abs(-2)
	fmt.Println(res)

	res = math.Abs(2)
	fmt.Println(res)

	res = math.Abs(3.5)
	fmt.Println(res)

	res = math.Abs(-3.5)
	fmt.Println(res)
}
