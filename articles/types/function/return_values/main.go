package main

import "fmt"

func main() {
	sum, avg := sum_avg(4, 2)
	fmt.Println(sum)
	fmt.Println(avg)
}

func sum_avg(a, b int) (int, int) {
	sum := a + b
	avg := (a + b) / 2
	return sum, avg
}
