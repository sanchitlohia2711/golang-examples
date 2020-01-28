package main

import "fmt"

func main() {

	eventsChan := make(chan string)

	go sendEvents(eventsChan)

	for event := range eventsChan {
		fmt.Println(event)
	}
}

func sendEvents(eventsChan chan<- string) {
	eventsChan <- "a"
	eventsChan <- "b"
	eventsChan <- "c"
	close(eventsChan)
}
