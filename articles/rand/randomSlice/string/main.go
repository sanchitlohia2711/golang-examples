package main

import (
	"fmt"
	"math/rand"
)

func main() {
	in := "abcdedf£"

	inRune := []rune(in)
	randomIndex := rand.Intn(len(inRune))
	pick := inRune[randomIndex]
	fmt.Println(string(pick))
}
