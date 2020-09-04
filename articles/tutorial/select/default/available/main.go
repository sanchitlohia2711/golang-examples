package main

import "fmt"

func main() {
	ch1 := make(chan string, 1)

	ch1 <- "Some value"
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	default:
		fmt.Println("Default statement executed")
	}
}
