package main

import "fmt"

func main() {
	s := test()
	fmt.Println(s)
}

func test() (size int) {
	defer func() { size = 20 }()
	size = 30
	return
}
