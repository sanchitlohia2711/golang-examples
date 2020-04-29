package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Provide seed
	rand.Seed(time.Now().Unix())

	//Generate a random array of length n
	fmt.Println(rand.Perm(10))

	fmt.Println(rand.Perm(10))

}
