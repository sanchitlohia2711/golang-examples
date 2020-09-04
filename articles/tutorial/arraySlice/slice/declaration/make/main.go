package main

import "fmt"

func main() {

	numbers := make([]int, 3, 5)

	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//With capacity ommited
	numbers = make([]int, 3)
	fmt.Println("\nWith Capacity Ommited")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

}
