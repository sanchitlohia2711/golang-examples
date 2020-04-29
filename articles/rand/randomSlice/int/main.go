package main

import (
	"fmt"
	"math/rand"
)

func main() {
	in := []int{2, 3, 5, 8}

	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]
	fmt.Println(pick)
}
