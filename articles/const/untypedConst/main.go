package main

import (
	"fmt"
	"math"
)

func main() {
	var f1 float32
	var f2 float64

	f1 = math.Pi
	f2 = math.Pi

	fmt.Printf("Type: %T Value: %v\n", math.Pi, math.Pi)
	fmt.Printf("Type: %T Value: %v\n", f1, f1)
	fmt.Printf("Type: %T Value: %v\n", f2, f2)
}
