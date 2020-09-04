package main

import "fmt"

func main() {
	sample := make([][]int, 2)
	sample[0] = []int{1, 2, 3}
	sample[1] = []int{4, 5, 6}

	fmt.Println("Using for-range")
	for _, row := range sample {
		for _, val := range row {
			fmt.Println(val)
		}
	}

	fmt.Println("\nUsing for loop - Second way")
	for i := 0; i < len(sample); i++ {
		for j := 0; j < len(sample[i]); j++ {
			fmt.Println(sample[i][j])
		}

	}
}
