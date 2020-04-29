package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Logb(4)
	fmt.Println(res)

	res = math.Logb(10.2)
	fmt.Println(res)

	res = math.Logb(-10)
	fmt.Println(res)
}
