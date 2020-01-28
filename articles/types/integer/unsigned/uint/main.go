package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main() {
	//This is computed as const uintSize = 32 << (^uuint(0) >> 32 & 1) // 32 or 64
	sizeOfuintInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfuintInBits)
	var a uint
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
