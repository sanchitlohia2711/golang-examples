package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	//Only lowercase
	charSet := []rune("abcdedfghijklmnopqrst")

	var output strings.Builder
	length := 10
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteRune(randomChar)
	}

	fmt.Println(output.String())
	output.Reset()

	//Lowercase and Uppercase Both
	charSet = []rune("abcdedfghijklmnopqrstABCDEFGHIJKLMNOP")
	length = 20
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteRune(randomChar)
	}
	fmt.Println(output.String())
}
