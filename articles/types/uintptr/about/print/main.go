package main

import (
	"fmt"
	"unsafe"
)

type sample struct {
	a int
	b string
}

func main() {
	s := &sample{
		a: 1,
		b: "test",
	}

	//Get the address as a uintptr
	startAddress := uintptr(unsafe.Pointer(s))
	fmt.Printf("Start Address of s: %d\n", startAddress)
}
