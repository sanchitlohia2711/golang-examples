package main

import "fmt"

func main() {

	sample := "Hello"

	suffix := "World"

	result := append([]byte(sample), suffix...)

	fmt.Printf("sample: %s\n", string(result))
}
