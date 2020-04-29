package main

import "fmt"

func main() {
	sample := "ab£c"
	sampleRune := []rune(sample)

	fmt.Printf("%c\n", sampleRune[0])
	fmt.Printf("%c\n", sampleRune[1])
	fmt.Printf("%c\n", sampleRune[2])
	fmt.Printf("%c\n", sampleRune[3])
}
