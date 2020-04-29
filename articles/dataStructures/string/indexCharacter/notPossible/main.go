package main

import "fmt"

func main() {
	sample := "ab£c"

	for i := 0; i < 4; i++ {
		fmt.Printf("%c\n", sample[i])
	}

	fmt.Printf("Length is %d\n", len(sample))
}
