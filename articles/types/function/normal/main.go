package main

import "fmt"

func main() {
	res := sum(2, 3)
	fmt.Println(res)
}

func sum(a, b int) int {
	return a + b
}
