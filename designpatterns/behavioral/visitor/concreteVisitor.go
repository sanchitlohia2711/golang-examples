package main

import "fmt"

type concreteVisitor struct {
}

func (c *concreteVisitor) visitAlpha(e element) {
	fmt.Println("Visit Logic Added for Alpha")
}

func (c *concreteVisitor) visitBeta(e element) {
	fmt.Println("Visit Logic Added for Beta")
}
