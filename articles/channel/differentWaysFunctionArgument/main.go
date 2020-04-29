package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	both(ch)
	onlySend(ch)

}

func both(ch chan int) {
	ch <- 2
	s := <-ch
	fmt.Println(s)
}

func onlySend(ch chan int) {
	ch <- 2
	s := <-ch
	fmt.Println(s)
}
