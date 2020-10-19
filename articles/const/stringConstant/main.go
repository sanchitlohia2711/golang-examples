package main

import "fmt"

func main() {
	type myString string

	//Typed String constant
	const aa string = "abc"
	var uu = aa
	fmt.Println("Untyped named string constant")
	fmt.Printf("uu: Type: %T Value: %v\n\nn", uu, uu)

	//Below line will raise a compilation error
	//var v myString = aa

	//Untyped named string constant
	const bb = "abc"
	var ww myString = bb
	var xx = bb
	fmt.Println("Untyped named string constant")
	fmt.Printf("ww: Type: %T Value: %v\n", ww, ww)
	fmt.Printf("xx: Type: %T Value: %v\n\n", xx, xx)

	//Untyped unnamed string constant
	var yy myString = "abc"
	var zz = "abc"
	fmt.Println("Untyped unnamed string constant")
	fmt.Printf("yy: Type: %T Value: %v\n", yy, yy)
	fmt.Printf("zz: Type: %T Value: %v\n", zz, zz)

}
