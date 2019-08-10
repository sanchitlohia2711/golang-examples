package main

type ball struct {
	radius int
	color  string
}

func (b *ball) clone() sportsKit {
	return &ball{
		radius: b.radius,
		color:  b.color,
	}
}
