package main

func main() {
	alpha := &alpha{}
	beta := &beta{}

	concreteVisitor := &concreteVisitor{}

	alpha.visit(concreteVisitor)
	beta.visit(concreteVisitor)
}
