package main

import "fmt"

func main() {
	sample := "ab£d"
	r := []rune(sample)
	var res []rune
	for i := len(r) - 1; i >= 0; i-- {
		res = append(res, r[i])
	}

	fmt.Printf("Result: %s\n", string(res))
}
