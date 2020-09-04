package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	// ch <- 2
	// ch <- 2
	// ch <- 2
	//close(ch)
	go sum(ch)
	time.Sleep(time.Second * 5)
}

func sum(ch chan int) {
	sum := 0
	for val := range ch {
		sum += val
	}
	fmt.Printf("Sum: %d\n", sum)
}
