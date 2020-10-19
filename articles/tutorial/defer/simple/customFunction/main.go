package main

import "fmt"

func main() {
	defer test()
	fmt.Println("Executed")
}

func test() {
	fmt.Println("In Defer")
}
