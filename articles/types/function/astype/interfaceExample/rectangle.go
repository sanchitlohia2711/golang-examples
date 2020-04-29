package main

type rectangle struct {
	length int
	breath int
}

func (r *rectangle) area() int {
	return r.length * r.breath
}

func (r *rectangle) getType() string {
	return "rectangle"
}
