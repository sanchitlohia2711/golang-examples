package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.RoundToEven(0.5)
	fmt.Println(res)

	res = math.RoundToEven(1.5)
	fmt.Println(res)

	res = math.RoundToEven(2.5)
	fmt.Println(res)

	res = math.RoundToEven(3.5)
	fmt.Println(res)

	res = math.RoundToEven(4.5)
	fmt.Println(res)
}
