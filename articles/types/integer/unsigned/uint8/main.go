package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a uint8
	var a uint8 = 2

	//Size of uint8 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
