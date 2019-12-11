package main

import "fmt"

func main() {

	handle(1, "abc")
	handle("abc", "xyz", 3)
	handle(1, 2, 3, 4)

}

func handle(args ...interface{}) {
	fmt.Println("Handle func called with parameters:")
	for _, arg := range args {
		fmt.Printf("%v\n", arg)
	}
}
