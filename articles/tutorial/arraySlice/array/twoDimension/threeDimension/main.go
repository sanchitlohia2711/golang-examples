package main

import "fmt"

func main() {
	sample := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}

	fmt.Printf("Length of first dimension: %d\n", len(sample))
	fmt.Printf("Length of second dimension: %d\n", len(sample[0]))
	fmt.Printf("Length of third dimension: %d\n", len(sample[0][0]))

	fmt.Printf("Overall Dimension of the array: %d*%d*%d\n", len(sample), len(sample[0]), len(sample[0][0]))
	fmt.Printf("Total number of elements in array: %d\n", len(sample)*len(sample[0])*len(sample[0][0]))

	for _, first := range sample {
		for _, second := range first {
			for _, value := range second {
				fmt.Println(value)
			}
		}
	}

}
