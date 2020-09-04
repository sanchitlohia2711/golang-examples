package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	ch <- 2
	ch <- 2
	ch <- 2
	close(ch)
	sum(ch)
	time.Sleep(time.Second * 1)
}

func sum(ch chan int) {
	sum := 0
	for i := 0; i < 2; i++ {
		val := <-ch
		sum += val
	}
	fmt.Printf("Sum: %d\n", sum)
}
