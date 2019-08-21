package main

import "fmt"

type A interface {
	walk()
	talk()
}
type B struct {
	name string
}

func (b B) walk() {
	fmt.Println("Work")
}
func (b B) talk() {
	fmt.Println("Talk")
}

func main() {
	b := 1
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", b)
}

func test(a A) {
	a.walk()
	a.talk()
}
