package main

import "fmt"

func main() {

	s := make([]int, 3)
	for _, v := range s {
		fmt.Println(v)
	}

	fmt.Println(len(s), cap(s))

	s = append(s, 2)
	for _, v := range s {
		fmt.Println(v)
	}

	fmt.Println(len(s), cap(s))

}
