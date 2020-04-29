package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Round(1.6)
	fmt.Println(res)

	res = math.Round(1.5)
	fmt.Println(res)

	res = math.Round(1.4)
	fmt.Println(res)

	res = math.Round(-1.6)
	fmt.Println(res)

	res = math.Round(-1.5)
	fmt.Println(res)

	res = math.Round(-1.4)
	fmt.Println(res)

	res = math.Round(1)
	fmt.Println(res)
}
