package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Log(4)
	fmt.Println(res)

	res = math.Log(10.2)
	fmt.Println(res)

	res = math.Log(-10)
	fmt.Println(res)
}
