package main

import "fmt"

func main() {
	twoDSlice := make([][]int, 2)
	twoDSlice[0] = []int{1, 2, 3}
	twoDSlice[1] = []int{4, 5}

	fmt.Printf("Number of rows in slice: %d\n", len(twoDSlice))
	fmt.Printf("Len of first row: %d\n", len(twoDSlice[0]))
	fmt.Printf("Len of second row: %d\n", len(twoDSlice[1]))

	fmt.Println("Traversing slice")
	for _, row := range twoDSlice {
		for _, val := range row {
			fmt.Println(val)
		}
	}

}
