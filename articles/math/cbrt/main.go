package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Cbrt(8)
	fmt.Println(res)

	res = math.Cbrt(27)
	fmt.Println(res)

	res = math.Cbrt(30.33)
	fmt.Println(res)
}
