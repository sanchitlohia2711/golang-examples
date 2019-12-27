package main

import "fmt"

func main() {

	if 0.1+0.2 == 0.3 {
		fmt.Println("Same")
	} else {
		fmt.Println("Not Same")
	}

	a := 3.14
	b := 3.14
	if a == b {
		fmt.Println("Same")
	} else {
		fmt.Println("Not Same")
	}
	a = 3.142
	b = 3.14
	if a == b {
		fmt.Println("Same")
	} else {
		fmt.Println("Not Same")
	}
}
