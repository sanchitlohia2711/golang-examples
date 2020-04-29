package main

import (
	"fmt"
	"math"
)

func main() {
	max := math.Max(2, 3)
	fmt.Println(max)

	max = math.Max(-2.1, -3.3)
	fmt.Println(max)
}
