package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Mod(4, 2)
	fmt.Println(res)

	res = math.Mod(4.2, 2)
	fmt.Println(res)

	res = math.Mod(5, 2)
	fmt.Println(res)

	res = math.Mod(-5, 2)
	fmt.Println(res)
}
