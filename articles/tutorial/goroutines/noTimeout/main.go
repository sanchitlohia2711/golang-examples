package main

import (
	"fmt"
)

func main() {
	go start()
	fmt.Println("Started")
	fmt.Println("Finished")
}

func start() {
	fmt.Println("In Goroutine")
}
