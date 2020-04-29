package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Log2(4)
	fmt.Println(res)

	res = math.Log2(10)
	fmt.Println(res)

	res = math.Log2(-10)
	fmt.Println(res)

}
