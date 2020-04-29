package main

import "fmt"

type area func(int, int) int

func main() {

	var areaF2 area = func(a, b int) int {
		return a * b
	}
	print(2, 3, areaF2)
}

func print(x, y int, a area) {
	fmt.Printf("Area is: %d\n", a(x, y))
}
