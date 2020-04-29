package main

import "fmt"

func main() {
	sample := "ab£c"
	s := []rune(sample)
	res := delChar(s, 2)
	fmt.Println(string(res))
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}
