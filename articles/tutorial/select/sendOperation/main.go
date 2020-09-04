package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go goOne(ch1)
	go goTwo(ch2)

	select {
	case ch1 <- "To goOne goroutine":
	case ch2 <- "To goTwo goroutine":
	}
	time.Sleep(time.Second * 1)
}

func goOne(ch chan string) {
	msg := <-ch
	fmt.Println(msg)
}

func goTwo(ch chan string) {
	msg := <-ch
	fmt.Println(msg)
}
