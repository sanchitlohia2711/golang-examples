package main

import "fmt"

type cashier struct {
}

func (c *cashier) execute(*patient) {
	fmt.Println("Cashier getting money from patient patient")
}
