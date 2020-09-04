package main

import (
	"fmt"
	"sync"
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

func merge(cs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	output := func(c <-chan string) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	chars := []string{"a", "b", "c"}
	done := make(chan struct{})

	in := startPipeline(done, chars)

	c1 := append(done, in)
	c2 := append(done, in)

	c3 := merge(c1, c2)
	endPipeline(done, c3)

	// fmt.Println(<-done)
	// fmt.Println(<-done)
	// fmt.Println(<-done)
	// fmt.Println(<-done)
	time.Sleep(4 * time.Second)
}
