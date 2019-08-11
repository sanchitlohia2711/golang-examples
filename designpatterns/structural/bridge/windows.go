package main

import "fmt"

type window struct {
	printer printer
}

func (m *window) print() {
	m.printer.printFile()
}

func (m *window) display() {
	fmt.Println("Display on a window monitor")
}
