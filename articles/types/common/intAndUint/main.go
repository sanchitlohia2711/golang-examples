package main

import (
	"fmt"
	"math/bits"
	"unsafe"
)

func main() {
	//This is computed as
	//const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	sizeInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeInBits)

	//Using unsafe.Sizeof() function. It will print size in bytes
	var a int
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))

	//Using unsafe.Sizeof() function. It will print size in bytes
	var b uint
	fmt.Printf("%d bytes\n", unsafe.Sizeof(b))
}
