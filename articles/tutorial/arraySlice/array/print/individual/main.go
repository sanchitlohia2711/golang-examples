package main

import "fmt"

func main() {
	numbers := [3]int{1, 2, 3}

	for _, num := range numbers {
		fmt.Println(num)
	}
}
