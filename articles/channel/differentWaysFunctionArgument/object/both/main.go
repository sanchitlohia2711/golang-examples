package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	process(ch)
}

func process(ch chan int) {
	ch <- 2
	s := <-ch
	fmt.Println(s)
}
