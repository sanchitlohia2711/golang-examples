package main

import (
	"fmt"
	"math"
)

func main() {
	//Power for integers
	res := math.Pow(2, 10)
	fmt.Println(res)

	//Power for float
	res = math.Pow(1.5, 2)
	fmt.Println(res)

	//Anything to power 0 is 1
	res = math.Pow(3, 0)
	fmt.Println(res)
}
