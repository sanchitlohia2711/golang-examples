package main

import "fmt"

type docter struct {
	nextDepartment department
}

func (d *docter) execute(p *patient) {
	fmt.Println("Docter checking patient")
	d.nextDepartment.execute(p)
}
