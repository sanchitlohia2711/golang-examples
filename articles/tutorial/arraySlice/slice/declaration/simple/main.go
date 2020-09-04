package main

import "fmt"

func main() {

	sample := []int{}
	fmt.Println(len(sample))
	fmt.Println(cap(sample))
	fmt.Println(sample)

	letters := []string{"a", "b", "c"}
	fmt.Println(len(letters))
	fmt.Println(cap(letters))
	fmt.Println(letters)
}
