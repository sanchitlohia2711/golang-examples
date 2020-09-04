package main

import "fmt"

var aaa = "test"

func main() {
	testGlobal()
}

func testGlobal() {
	fmt.Println(aaa)
}
