package main

import "fmt"

func main() {
	sample := make([][]int, 2)
	sample[0] = []int{1, 2, 3}
	sample[1] = []int{4, 5, 6}

	//Print array element
	fmt.Println(sample[0][0])
	fmt.Println(sample[0][1])
	fmt.Println(sample[0][2])
	fmt.Println(sample[1][0])
	fmt.Println(sample[1][1])
	fmt.Println(sample[1][2])

	//Assign new values
	sample[0][0] = 6
	sample[0][1] = 5
	sample[0][2] = 4
	sample[1][0] = 3
	sample[1][1] = 2
	sample[1][2] = 1

	fmt.Println()
	fmt.Println(sample[0][0])
	fmt.Println(sample[0][1])
	fmt.Println(sample[0][2])
	fmt.Println(sample[1][0])
	fmt.Println(sample[1][1])
	fmt.Println(sample[1][2])
}
