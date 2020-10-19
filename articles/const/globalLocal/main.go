package main

import "fmt"

const name = "test"

func main() {
	val := data.PI
}

func testGlobal() {
	fmt.Println(name)
	//The below line will give compiler error
	//fmt.Println(a)
}
