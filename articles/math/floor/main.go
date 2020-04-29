package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Floor(1.6)
	fmt.Println(res)

	res = math.Floor(-1.6)
	fmt.Println(res)

	res = math.Floor(1)
	fmt.Println(res)
}
