package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	ch <- "Before break"

	select {
	case msg := <-ch:
		fmt.Println(msg)
		break
		fmt.Println("After break")
	default:
		fmt.Println("Default case")
	}
}
