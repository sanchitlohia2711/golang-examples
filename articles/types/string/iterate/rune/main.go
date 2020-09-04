package main

import "fmt"

func main() {
	sample := "a£b"

	runeSample := []rune(sample)

	fmt.Printf("Length of given string is %d\n", len(runeSample))

	//Iterate
	for i := 0; i < len(runeSample); i++ {
		fmt.Println(string(runeSample[i]))
	}
}
