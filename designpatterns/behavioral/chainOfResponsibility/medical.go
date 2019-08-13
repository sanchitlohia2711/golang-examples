package main

import "fmt"

type medical struct {
	nextDepartment department
}

func (m *medical) execute(p *patient) {
	fmt.Println("Medical giving medicine to patient")
	m.nextDepartment.execute(p)
}
