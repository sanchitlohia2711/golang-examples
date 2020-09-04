package main

import "fmt"

//Test function
func main() {
	//STRUCTURE IDENTIFIER
	p := &Person{
		Name: "test",
		age:  21,
	}
	fmt.Println(p)
	c := &company{}
	fmt.Println(c)

	//STRUCTURE'S FIELDS
	fmt.Println(p.Name)
	fmt.Println(p.age)

	//STRUCTURE'S METHOD
	fmt.Println(p.GetAge())
	fmt.Println(p.getName())
}
