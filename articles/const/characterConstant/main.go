package main

import "fmt"

func main() {
	type myChar int32

	//Typed character constant
	const aa int32 = 'a'
	var uu = aa
	fmt.Println("Untyped unnamed character constant")
	fmt.Printf("uu: Type: %T Value: %v\n\n", uu, uu)

	//Below line will raise a compilation error
	//var vv myBool = aa

	//Untyped named character constant
	const bb = 'a'

	var ww myChar = bb
	var xx = bb
	fmt.Println("Untyped named character constant")
	fmt.Printf("ww: Type: %T Value: %v\n", ww, ww)
	fmt.Printf("xx: Type: %T Value: %v\n\n", xx, xx)

	//Untyped unnamed character constant
	var yy myChar = 'a'
	var zz = 'a'
	fmt.Println("Untyped unnamed character constant")
	fmt.Printf("yy: Type: %T Value: %v\n", yy, yy)
	fmt.Printf("zz: Type: %T Value: %v\n", zz, zz)
}
