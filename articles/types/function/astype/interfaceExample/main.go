package main

import "fmt"

func main() {
	var shapes []shape

	s := &square{side: 2}
	shapes = append(shapes, s)

	r := &rectangle{length: 2, breath: 3}
	shapes = append(shapes, r)

	for _, shape := range shapes {
		fmt.Printf("Type: %s, Area %d\n", shape.getType(), shape.area())
	}
}
