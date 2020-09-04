package main

import "fmt"

func main() {
	sample := [2][3]byte{}

	//First row
	fmt.Println("First row")
	fmt.Println(&sample[0][0])
	fmt.Println(&sample[0][1])
	fmt.Println(&sample[0][2])

	fmt.Println("\nSecond row")
	fmt.Println(&sample[1][0])
	fmt.Println(&sample[1][1])
	fmt.Println(&sample[1][2])
}
