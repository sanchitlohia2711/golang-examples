package main

import "fmt"

type mac struct {
	printer printer
}

func (m *mac) print() {
	m.printer.printFile()
}

func (m *mac) display() {
	fmt.Println("Display on a mac monitor")
}
