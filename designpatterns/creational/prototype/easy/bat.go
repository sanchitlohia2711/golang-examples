package main

type bat struct {
	length int
	color  string
}

func (b *bat) clone() sportsKit {
	return &bat{
		length: b.length,
		color:  b.color,
	}
}
