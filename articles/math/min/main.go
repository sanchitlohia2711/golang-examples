package main

import (
	"fmt"
	"math"
)

func main() {
	min := math.Min(2, 3)
	fmt.Println(min)

	min = math.Min(-2.1, -3.3)
	fmt.Println(min)
}
