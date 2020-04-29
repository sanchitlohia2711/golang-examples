package main

import (
	"fmt"
	"math"
)

func main() {
	//Square root of a integer
	res := math.Sqrt(4)
	fmt.Println(res)

	//Square root of a integer
	res = math.Sqrt(9)
	fmt.Println(res)

	//Square Root of a float
	res = math.Sqrt(30.33)
	fmt.Println(res)

	//Square Root of a negative number
	res = math.Sqrt(-9)
	fmt.Println(res)
}
