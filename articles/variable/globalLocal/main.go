package main

import "fmt"

var aaa = "test"

func main() {
	var bbb = "test2"
	fmt.Println(bbb)
	testGlobal()
}

func testGlobal() {
	fmt.Println(aaa)
	//The below line will give compiler error
	//fmt.Println(bbb)
}
