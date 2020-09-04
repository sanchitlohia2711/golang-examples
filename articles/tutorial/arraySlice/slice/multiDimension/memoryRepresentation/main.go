package main

import "fmt"

func main() {
	sample := make([][]byte, 2)
	sample[0] = make([]byte, 3)

	//testVariable := "s"
	//fmt.Println(testVariable)
	sample[1] = make([]byte, 3)

	fmt.Println("First row")
	fmt.Println(&sample[0][0])
	fmt.Println(&sample[0][1])
	fmt.Println(&sample[0][2])

	fmt.Println("\nSecond row")
	fmt.Println(&sample[1][0])
	fmt.Println(&sample[1][1])
	fmt.Println(&sample[1][2])
}
