package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Log10(100)
	fmt.Println(res)

	res = math.Log10(10)
	fmt.Println(res)

	res = math.Log10(-10)
	fmt.Println(res)

}
