package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[:]

	//Modifying the slice
	slice[1] = 7
	fmt.Println("Modifying Slice")
	fmt.Printf("Array=%v\n", array)
	fmt.Printf("Slice=%v\n", slice)

	//Modifying the array. Would reflect back in slcie too
	array[1] = 2
	fmt.Println("\nModifying Underlying Array")
	fmt.Printf("Array=%v\n", array)
	fmt.Printf("Slice=%v\n", slice)
}
