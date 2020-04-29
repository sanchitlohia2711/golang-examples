package main

import (
	"fmt"
	"math"
)

func main() {
	//Will return false for positive
	res := math.Signbit(4)
	fmt.Println(res)

	//Will return true for negative
	res = math.Signbit(-4)
	fmt.Println(res)

	//Will return false for zerp
	res = math.Signbit(0)
	fmt.Println(res)

	//Will return false for positive float
	res = math.Signbit(-0)
	fmt.Println(res)

}
