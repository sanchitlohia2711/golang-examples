package main

import "fmt"

func main() {
	//Untyped
	var u = 123      //Default hidden type is int
	var v = "circle" //Default hidden type is string
	var w = 5.6      //Default hidden type is float64
	var x = true     //Default hidden type is bool
	var y = 'a'      //Default hidden type is rune
	var z = 3 + 5i   //Default hidden type is complex128

	fmt.Printf("Type: %T Value: %v\n", u, u)
	fmt.Printf("Type: %T Value: %v\n", v, v)
	fmt.Printf("Type: %T Value: %v\n", w, w)
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
