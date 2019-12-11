package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}, 4)

	for i := 0; i < 4; i++ {
		go runasync(done)
	}

	for i := 0; i < 4; i++ {
		<-done
	}
	close(done)

	fmt.Printf("Finish")
}

func runasync(done chan<- struct{}) {
	time.Sleep(1 * time.Second)
	done <- struct{}{}
	return
}
