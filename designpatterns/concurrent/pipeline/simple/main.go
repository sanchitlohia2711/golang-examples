package main

import "fmt"

func append(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for s := range in {
			out <- s + s[len(s)-1:]
		}
		close(out)
	}()

	return out
}

func startPipeline(chars []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, s := range chars {
			out <- s + s
		}
		close(out)
	}()

	return out
}

func endPipeline(in <-chan string) {
	for s := range in {
		fmt.Println(s)
	}
	return
}

func main() {
	chars := []string{"a", "b", "c"}
	endPipeline(append(append(startPipeline(chars))))
}
