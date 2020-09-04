package main

import "fmt"

func main() {
	ch1 := make(chan string)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	}
}
