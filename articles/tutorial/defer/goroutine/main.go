package main

import (
	"fmt"
	"time"
)

func main() {
	call()
	time.Sleep(time.Second * 2)
}

func call() {
	defer func() {
		go test()
	}()
	fmt.Println("In call function")
}

func test() {
	fmt.Println("In test function")
}
