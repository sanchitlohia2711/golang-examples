package main

import "fmt"

func main() {
	sample := "ab£c"

	fmt.Println([]byte(sample))

	fmt.Printf("Length is %d\n", len(sample))
}
