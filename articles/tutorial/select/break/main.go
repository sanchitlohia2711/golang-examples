package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	ch <- "Before break"

	select {
	case msg1 := <-ch:
		fmt.Println(msg1)
		break
		fmt.Println("After break")
	default:
		fmt.Println("Default case")
	}
}
