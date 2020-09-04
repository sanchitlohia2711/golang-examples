package main

import "fmt"

func main() {
	sample := [2]string{"aa", "bb"}

	fmt.Println(sample[0])
	fmt.Println(sample[1])

	sample[0] = "xx"

	fmt.Println(sample)

	//sample[3] = "yy"
}
