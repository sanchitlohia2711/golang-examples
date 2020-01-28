package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	r := '£'

	//Print Size
	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))

	//Print Type
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))

	//Print Code Point
	fmt.Printf("Unicode CodePoint: %U\n", r)

	//Print Character
	fmt.Printf("Character: %c\n", r)

	s := "0b£"

	//This will print the Unicode Points
	fmt.Printf("%U\n", []rune(s))

	//This will the decimal value of Unicode Code Point
	fmt.Println([]rune(s))
}
