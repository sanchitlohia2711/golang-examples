package main

import "fmt"

func main() {
	sample := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("First Run")
	for _, row := range sample {
		for _, val := range row {
			fmt.Println(val)
		}
	}

	sample[0][0] = 6
	sample[1][2] = 1

	fmt.Println("\nSecond Run")
	for _, row := range sample {
		for _, val := range row {
			fmt.Println(val)
		}
	}
}
