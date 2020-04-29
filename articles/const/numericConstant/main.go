package main

import "fmt"

func main() {
	//Typed int constant
	const aa int = 123
	var uu = aa
	fmt.Printf("uu: Type: %T Value: %v\n", uu, uu)

	//Below line will raise a compilation error
	//var v int32 = aa

	//Untyped named int constant
	const bb = 123
	var ww = bb
	var xx int32 = bb
	var yy float64 = bb
	var zz complex128 = bb
	fmt.Printf("ww: Type: %T Value: %v\n", ww, ww)
	fmt.Printf("xx: Type: %T Value: %v\n", xx, xx)
	fmt.Printf("yy: Type: %T Value: %v\n", yy, yy)
	fmt.Printf("zz: Type: %T Value: %v\n", zz, zz)

	//Untyped unnamed int constant
	const cc = 123
	var ll = cc
	var mm int32 = bb
	var nn float64 = bb
	var oo complex128 = bb
	fmt.Printf("ll: Type: %T Value: %v\n", ll, ll)
	fmt.Printf("mm: Type: %T Value: %v\n", mm, mm)
	fmt.Printf("nn: Type: %T Value: %v\n", nn, nn)
	fmt.Printf("oo: Type: %T Value: %v\n", oo, oo)
}
