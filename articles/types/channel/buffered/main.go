package main

import "fmt"

func main() {
	//Creating a buffered channel of length 3
	eventsChan := make(chan string, 3)

	eventsChan <- "a"
	eventsChan <- "b"
	eventsChan <- "c"
	//Closing the channel
	close(eventsChan)

	for event := range eventsChan {
		fmt.Println(event)
	}
}
