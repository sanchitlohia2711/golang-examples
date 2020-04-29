package main

import "fmt"

func main() {
	count := 0
	for i := 1; i <= 5; i++ {
		func() {
			count++
			fmt.Println(count)
		}()
	}
}
