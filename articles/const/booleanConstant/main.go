package main

import "fmt"

func main() {
	type myBool bool

	//Typed String constant
	const aa bool = true
	var uu = aa
	fmt.Printf("uu: Type: %T Value: %v\n", uu, uu)

	//Below line will raise a compilation error
	//var vv myBool = aa

	//Untyped named string constant
	const bb = true

	var ww myBool = bb
	var xx = bb
	fmt.Printf("ww: Type: %T Value: %v\n", ww, ww)
	fmt.Printf("xx: Type: %T Value: %v\n", xx, xx)

	//Untyped unnamed sting constant
	const cc = true
	var yy myBool = cc
	var zz = cc
	fmt.Printf("yy: Type: %T Value: %v\n", yy, yy)
	fmt.Printf("zz: Type: %T Value: %v\n", zz, zz)
}
