package main

import "fmt"

type iBase interface {
	say()
}

type child struct {
	style string
}

func (b *child) say() {
	fmt.Println(b.style)
}

func check(b iBase) {
	b.say()
}

func main() {
	child := &child{
		style: "somestyle",
	}
	child.say()
	check(child)
}
