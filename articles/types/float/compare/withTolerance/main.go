package main

import (
	"fmt"
	"math"
)

func main() {
	withTolerane(3.14, 3.141)

	withTolerane(3.14, 3.142)

}

func withTolerane(a, b float64) {
	tolerance := 0.001
	if diff := math.Abs(a - b); diff < tolerance {
		fmt.Printf("When a=%f and b =%f => Nearly same by tolerance\n", a, b)
	} else {
		fmt.Printf("When a=%f and b=%f => Not same Even by Tolerance\n", a, b)
	}
}
