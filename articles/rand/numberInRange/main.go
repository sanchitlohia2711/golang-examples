package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	//Generate a random number x where x is in range 5<=x<=20
	rangeLower := 5
	rangeUpper := 20
	randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
	fmt.Println(randomNum)

	//Generate a random number x where x is in range 100<=x<=200
	rangeLower = 100
	rangeUpper = 200
	randomNum = rangeLower + rand.Intn(rangeUpper-rangeLower+1)
	fmt.Println(randomNum)
}
