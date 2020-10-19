package main

import "fmt"

func main() {
	sample := "abc"

	defer fmt.Printf("In defer sample is: %s\n", sample)
	sample = "xyz"
	defer fmt.Printf("Latest value of sample is: %s\n", sample)
}
