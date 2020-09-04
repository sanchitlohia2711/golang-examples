package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int

	go send(ch)
	time.Sleep(time.Second * 1)
}

func send(ch chan int) {
	fmt.Println("Sending value to channnel start")
	ch <- 1
	fmt.Println("Sending value to channnel finish")
}
