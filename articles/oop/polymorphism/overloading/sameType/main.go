package main

import "fmt"

func main() {

	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2, 3, 4))

}

func add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
