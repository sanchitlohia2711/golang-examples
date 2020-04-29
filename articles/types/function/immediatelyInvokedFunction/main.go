package main

import "fmt"

func main() {
	squareOf2 := func() int {
		return 2 * 2
	}()
	fmt.Println(squareOf2)
}
