package main

import (
	"fmt"
	"time"
)

func forever(id int) {
	fmt.Printf("id: %d\n", id)
	for {
	}
}
func main() {

	for i := 0; i < 100; i++ {
		go forever(i)
	}
	time.Sleep(time.Minute * 1)
}
