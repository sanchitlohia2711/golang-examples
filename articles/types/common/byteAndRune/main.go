package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//If you don't specify type here
	var b byte = 'a'

	fmt.Println("Priting Byte:")
	//Print Size, Type and Character
	fmt.Printf("Size: %d\nType: %s\nCharacter: %c\n", unsafe.Sizeof(b), reflect.TypeOf(b), b)

	r := '£'

	fmt.Println("\nPriting Rune:")
	//Print Size, Type, CodePoint and Character
	fmt.Printf("Size: %d\nType: %s\nUnicode CodePoint: %U\nCharacter: %c\n", unsafe.Sizeof(r), reflect.TypeOf(r), r, r)

	s := "µ" //Micrs sign
	fmt.Println("\nPriting String:")
	fmt.Printf("Size: %d\nType: %s\nCharacter: %s\n", unsafe.Sizeof(s), reflect.TypeOf(s), s)
}
