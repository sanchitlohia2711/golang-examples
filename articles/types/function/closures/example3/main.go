package main

import "fmt"

func main() {
	valueOutside := "somevalue"
	func() {
		fmt.Println(valueOutside)
	}()
}
