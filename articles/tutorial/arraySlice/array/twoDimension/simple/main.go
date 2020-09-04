package main

import "fmt"

func main() {
	sample := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Printf("Number of rows in array: %d\n", len(sample))
	fmt.Printf("Number of columns in array: %d\n", len(sample[0]))
	fmt.Printf("Total number of elements in array: %d\n", len(sample)*len(sample[0]))
	fmt.Printf("Len of first row: %d\n", len(sample[0]))
	fmt.Printf("Len of second row: %d\n", len(sample[1]))

	fmt.Println("Traversing Array")
	for _, row := range sample {
		for _, val := range row {
			fmt.Println(val)
		}
	}

}
