package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	fmt.Println("Sending value to channnel start")
	go send(ch)

	val := <-ch
	fmt.Printf("Receiving Value from channel finished. Value received: %d\n", val)

}

func send(ch chan int) {
	ch <- 1
}
