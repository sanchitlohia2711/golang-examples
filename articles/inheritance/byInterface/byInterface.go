package main

import "fmt"

type iBase interface {
	say()
}

type base struct {
	value string
}

func (b *base) say() {
	fmt.Println(b.value)
}

type child struct {
	base  //embedding
	style string
}

func check(b iBase) {
	b.say()
}

func main() {

	base := base{value: "somevalue"}
	child := &child{
		base:  base,
		style: "somestyle",
	}

	child.say()

	check(child)
}
