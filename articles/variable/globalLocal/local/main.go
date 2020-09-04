package main

import "fmt"

func main() {
	var aaa = "test"
	fmt.Println(aaa)

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)
}

func testLocal() {
	fmt.Println(aaa)
}
