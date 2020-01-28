package main

import "fmt"

func main() {
	s := "ab£"

	r := []rune(s)

	fmt.Printf("%U\n", r)
}
