package main

import "fmt"

func main() {
	numbers := new([]int)

	fmt.Printf("numbers=%v\n", *numbers)
	fmt.Printf("length=%d\n", len(*numbers))
	fmt.Printf("capacity=%d\n", cap(*numbers))

}
