package main

import "fmt"

func main() {
	sample := "ab£c"
	s := []rune(sample)

	fmt.Println(len(s))

	len := 0
	for range sample {
		len++
	}

	fmt.Println(len)
}
