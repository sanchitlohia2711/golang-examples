package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2))
}

func doOperation(fn func(int, int) int, x, y int) int {
	return fn(x, y)
}
