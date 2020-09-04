package main

import "fmt"

func main() {
	numbers1 := []int{1, 2}

	numbers2 := []int{3, 4}

	numbers := append(numbers1, numbers2...)

	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}
