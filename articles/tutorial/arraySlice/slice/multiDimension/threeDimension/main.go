package main

import "fmt"

func main() {
	sample := make([][][]int, 2)
	for i := range sample {
		sample[i] = make([][]int, 2)
		for j := range sample[i] {
			sample[i][j] = make([]int, 3)
		}
	}

	fmt.Printf("Length of first dimension: %d\n", len(sample))
	fmt.Printf("Length of second dimension: %d\n", len(sample[0]))
	fmt.Printf("Length of third dimension: %d\n", len(sample[0][0]))
	fmt.Printf("Overall Dimension of the slice: %d*%d*%d\n", len(sample), len(sample[0]), len(sample[0][0]))
	fmt.Printf("Total number of elements in slice: %d\n", len(sample)*len(sample[0])*len(sample[0][0]))

	for _, first := range sample {
		for _, second := range first {
			for _, value := range second {
				fmt.Println(value)
			}
		}
	}

}
