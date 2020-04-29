package main

import "fmt"

func(){
	temperature := 32
	fmt.Println(temperature)
}()

func main() {
	squareOf2 := func() int {
		res := 2 * 2
		return res
	}()
	fmt.Println(squareOf2)
}
