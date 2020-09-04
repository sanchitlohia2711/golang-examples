package main

import (
	"fmt"
	"time"
)

func append(done chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			fmt.Printf("append %s\n", s)
			select {
			case out <- s + s[len(s)-1:]:
			case <-done:
				fmt.Println("Done is closed in append")
				return
			}
		}
	}()

	return out
}

func startPipeline(done chan struct{}, chars []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, s := range chars {
			fmt.Printf("startPipiline %s\n", s)
			select {
			case out <- s + s:
			case <-done:
				fmt.Println("Done is closed in startPipeline")
				return
			}
		}

	}()

	return out
}

func endPipeline(done chan struct{}, in <-chan string) {
	defer close(done)
	fmt.Printf("endPipeline %s\n", <-in)
	fmt.Printf("endPipeline %s\n", <-in)
	fmt.Printf("endPipeline %s\n", <-in)

	//for s := range in {
	//	fmt.Println(s)
	//}
	return
}

func main() {
	chars := []string{"a", "b", "c"}
	done := make(chan struct{})
	endPipeline(done, append(done, startPipeline(done, chars)))

	// fmt.Println(<-done)
	// fmt.Println(<-done)
	// fmt.Println(<-done)
	// fmt.Println(<-done)
	time.Sleep(4 * time.Second)
}
