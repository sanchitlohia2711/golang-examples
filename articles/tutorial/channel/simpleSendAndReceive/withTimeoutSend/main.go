package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second * 2)
}

func send(ch chan int) {
	ch <- 1
	fmt.Println("Sending value to channnel complete")
}

func receive(ch chan int) {
	time.Sleep(time.Second * 1)
	fmt.Println("Timeout finished")
	_ = <-ch
	return
}
