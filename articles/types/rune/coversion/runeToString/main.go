package main

import "fmt"

func main() {
	runeArray := []rune{'a', 'b', '£'}

	s := string(runeArray)
	fmt.Println(s)
}
