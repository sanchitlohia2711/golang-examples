package main

import "fmt"

func main() {
	sample := "ab£d"
	r := []rune(sample)

	fmt.Printf("Before: %s\n", string(r))
	r[2], r[3] = r[3], r[2]
	fmt.Printf("After: %s\n", string(r))
}
