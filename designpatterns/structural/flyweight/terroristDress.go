package main

type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}
