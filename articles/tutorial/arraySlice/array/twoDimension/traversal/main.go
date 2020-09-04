package main

import "fmt"

func main() {
	sample := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("Using for-range")
	for _, row := range sample {
		for _, val := range row {
			fmt.Println(val)
		}
	}

	fmt.Println("\nUsing for loop")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(sample[i][j])
		}
	}

	fmt.Println("\nUsing for loop - Second way")
	for i := 0; i < len(sample); i++ {
		for j := 0; j < len(sample[i]); j++ {
			fmt.Println(sample[i][j])
		}

	}
}
