package main

import "fmt"

func main() {
	areaF := getAreaFunc()
	res := areaF(2, 4)
	fmt.Println(res)
}

func getAreaFunc() func(int, int) int {
	return func(x, y int) int {
		return x * y
	}
}
