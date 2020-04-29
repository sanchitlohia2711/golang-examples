package main

import (
	"fmt"
	"math"
)

func main() {
	//Contain both integer and fraction
	integer, fraction := math.Modf(2.5)
	fmt.Printf("Integer: %f. Fraction: %f\n", integer, fraction)

	//Contain only integer part
	integer, fraction = math.Modf(2)
	fmt.Printf("Integer: %f. Fraction: %f\n", integer, fraction)

	//Negative floating point number
	integer, fraction = math.Modf(-2.5)
	fmt.Printf("Integer: %f. Fraction: %f\n", integer, fraction)

	//Contains only fraction part
	integer, fraction = math.Modf(0.5)
	fmt.Printf("Integer: %f. Fraction: %f\n", integer, fraction)
}
