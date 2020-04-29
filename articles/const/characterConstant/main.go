package main

import "fmt"

func main() {
	type myChar int32

	//Typed character constant
	const aa int32 = 'a'
	var uu = aa
	fmt.Printf("uu: Type: %T Value: %v\n", uu, uu)

	//Below line will raise a compilation error
	//var vv myBool = aa

	//Untyped named character constant
	const bb = 'a'

	var ww myChar = bb
	var xx = bb
	fmt.Printf("ww: Type: %T Value: %v\n", ww, ww)
	fmt.Printf("xx: Type: %T Value: %v\n", xx, xx)

	//Untyped unnamed character constant
	const cc = 'a'
	var yy myChar = cc
	var zz = cc
	fmt.Printf("yy: Type: %T Value: %v\n", yy, yy)
	fmt.Printf("zz: Type: %T Value: %v\n", zz, zz)
}
