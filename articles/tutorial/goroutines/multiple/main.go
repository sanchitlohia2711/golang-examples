package main

import (
	"fmt"
	"time"
)

func forever(id int) {
	fmt.Printf("id: %d\n", id)
}
func main() {
	fmt.Println("Started")
	for i := 0; i < 10; i++ {
		go forever(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("Finished")
}
