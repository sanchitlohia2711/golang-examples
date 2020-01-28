package main

import "fmt"

func main() {

	test("astring")
	test("10")
	test(true)
}

func test(a interface{}) {
	fmt.Printf("(%v, %T)\n", a, a)
}
