package main

import (
	"context"
	"fmt"
	"time"
)

func append(ctx context.Context, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			fmt.Printf("append %s\n", s)
			select {
			case out <- s + s[len(s)-1:]:
			case <-ctx.Done():
				fmt.Println("Done is closed in append")
				return
			}
		}
	}()

	return out
}

func startPipeline(ctx context.Context, chars []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, s := range chars {
			fmt.Printf("startPipiline %s\n", s)
			select {
			case out <- s + s:
			case <-ctx.Done():
				fmt.Println("Done is closed in startPipeline")
				return
			}
		}

	}()

	return out
}

func endPipeline(ctx context.Context, cancel func(), in <-chan string) {
	defer cancel()
	fmt.Printf("endPipeline %s\n", <-in)
	//fmt.Printf("endPipeline %s\n", <-in)
	//fmt.Printf("endPipeline %s\n", <-in)

	//for s := range in {
	//	fmt.Println(s)
	//}
	return
}

func main() {
	chars := []string{"a", "b", "c"}
	c := context.Background()
	ctx, cancel := context.WithCancel(c)
	endPipeline(ctx, cancel, append(ctx, startPipeline(ctx, chars)))

	// fmt.Println(<-done)
	// fmt.Println(<-done)
	// fmt.Println(<-done)
	// fmt.Println(<-done)
	time.Sleep(4 * time.Second)
}
