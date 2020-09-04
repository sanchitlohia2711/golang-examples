package main

import "fmt"

func main() {
	printType("test_string")
	printType(2)
	sppu := sp{}
	fmt.Printf("ww: Type: %T Value: %v\n", sppu, sppu)
}

func printType(t interface{}) {
	fmt.Printf("ww: Type: %T Value: %v\n", t, t)
	switch v := t.(type) {
	case string:
		fmt.Println("Type: string")
	case int:
		fmt.Println("Type: int")
	default:
		fmt.Printf("Unknown Type %T", v)
	}
}

type s interface {
	p()
}

type sp struct {
	name string
}

func (sa *sp) p() {
	return
}
