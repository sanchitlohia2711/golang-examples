package main

import "fmt"

func main() {
	t := 123      //Type Inferred will be int
	u := "circle" //Type Inferred will be string
	v := 5.6      //Type Inferred will be float64
	w := true     //Type Inferred will be bool
	x := 'a'      //Type Inferred will be rune
	y := 3 + 5i   //Type Inferred will be complex128
	z := &a{name: "test"}

	fmt.Printf("Type: %T Value: %v\n", t, t)
	fmt.Printf("Type: %T Value: %v\n", u, u)
	fmt.Printf("Type: %T Value: %v\n", v, v)
	fmt.Printf("Type: %T Value: %v\n", w, w)
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", z, z)

}

type a struct {
	name string
}
