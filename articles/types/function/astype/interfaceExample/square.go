package main

type square struct {
	side int
}

func (s *square) area() int {
	return s.side * s.side
}

func (s *square) getType() string {
	return "square"
}
