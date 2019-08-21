package main

type alpha struct {
}

func (a *alpha) visit(v visitor) {
	v.visitAlpha(a)
}
