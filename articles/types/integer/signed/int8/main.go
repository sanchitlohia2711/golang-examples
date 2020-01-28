package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a int8 = 2
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
