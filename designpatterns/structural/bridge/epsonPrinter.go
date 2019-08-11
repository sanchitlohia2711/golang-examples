package main

import "fmt"

type epsonPrinter struct {
}

func (p *epsonPrinter) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}
