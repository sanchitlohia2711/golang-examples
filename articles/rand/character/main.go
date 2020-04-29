package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	//Generate a random character between lowercase a to z
	randomChar := 'a' + rune(rand.Intn(26))
	fmt.Println(string(randomChar))

	//Generate a random character between upppercase A and Z
	randomChar = 'A' + rune(rand.Intn(26))
	fmt.Println(string(randomChar))

	//Generate a random character between uppercase A and Z  and lowercase a to z
	randomInt := rand.Intn(2)
	if randomInt == 1 {
		randomChar = 'A' + rune(rand.Intn(26))
	} else {
		randomChar = 'a' + rune(rand.Intn(26))
	}
	fmt.Println(string(randomChar))

}
