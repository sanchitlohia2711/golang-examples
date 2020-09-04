package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan int, 1)
	sum(2, 3, result)

	value := <-result
	fmt.Printf("Value: %d\n", value)
	close(result)

}

func sum(a, b int, result chan int) {
	sumValue := a + b
	time.Sleep(time.Second * 2)
	result <- sumValue
	return
}
