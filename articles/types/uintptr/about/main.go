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
	fmt.Printf("Offset of B:%d\n", startAddress)
	//Getting the address of b in struct s
	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))
	//Typecasting it to a string pointer and printing the value of it
	fmt.Println(*(*string)(p))

}
