package main

import "fmt"

type hpPrinter struct {
}

func (p *hpPrinter) printFile() {
	fmt.Println("Printing by a HP Printer")
}
