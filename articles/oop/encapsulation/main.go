package main

import (
	"fmt"

	"github.com/golang-examples/articles/oop/encapsulation/model"
)

func main() {
	//STRUCTURE IDENTIFIER
	p := &model.Person{
		Name: "test",
		age:  21,
	}
	fmt.Println(p)

	c := &model.company{}
	fmt.Println(c)

	//STRUCTURE'S METHOD
	fmt.Println(p.GetAge())
	fmt.Println(p.getName())

	//STRUCTURE'S FIELDS
	fmt.Println(p.Name)
	fmt.Println(p.age)

	//FUNCTION
	person2 := model.GetPerson()
	fmt.Println(person2)
	companyName := model.getCompanyName()
	fmt.Println(companyName)

	//VARIBLES
	fmt.Println(model.companyLocation)
	fmt.Println(model.CompanyName)
}
