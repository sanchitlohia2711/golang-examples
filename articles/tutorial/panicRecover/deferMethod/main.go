package main

import (
	"errors"
	"fmt"
)

type employee struct {
	name string
}

func (e *employee) setName(name string) error {
	defer e.setDefaultName()
	if len(name) < 3 {
		fmt.Println("Length of name passed is less than 3")
		return errors.New("Length of name cannnot be less than 3")
	}
	e.name = name
	return nil
}

func (e *employee) setDefaultName() {
	fmt.Println("In the setDefaultName function")
	if e.name == "" {
		e.name = "DefaultName"
		fmt.Println("Default name is set")
	}
}

func main() {
	e1 := &employee{}
	e1.setName("John")
	fmt.Printf("First employee name is: %s\n", e1.name)

	fmt.Println()
	e2 := &employee{}
	e2.setName("Ko")
	fmt.Printf("Second employee name is: %s\n", e2.name)
	return
}
