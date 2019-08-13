package main

import "fmt"

type reception struct {
	nextDepartment department
}

func (r *reception) execute(p *patient) {
	fmt.Println("Reception registering patient")
	r.nextDepartment.execute(p)
}
