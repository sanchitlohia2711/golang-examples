package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Remainder(4, 2)
	fmt.Println(res)

	res = math.Remainder(5, 2)
	fmt.Println(res)

	res = math.Remainder(5.5, 2)
	fmt.Println(res)

	res = math.Remainder(5.5, 1.5)
	fmt.Println(res)
}
