package main

import "fmt"

type folder struct {
	compositeArray []composite
	name           string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, composite := range f.compositeArray {
		composite.print(indentation + indentation)
	}
}
