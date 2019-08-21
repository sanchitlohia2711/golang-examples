package main

type beta struct {
}

func (b *beta) visit(v visitor) {
	v.visitBeta(b)
}
