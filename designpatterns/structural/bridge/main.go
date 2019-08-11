package main

func main() {

	hpPrinter := &hpPrinter{}
	epsonPrinter := &epsonPrinter{}

	macComputer1 := &mac{
		printer: hpPrinter,
	}

	macComputer1.print()

	macComputer2 := &mac{
		printer: epsonPrinter,
	}

	macComputer2.print()

	winComputer1 := &mac{
		printer: hpPrinter,
	}
	winComputer1.print()

	winComputer2 := &mac{
		printer: epsonPrinter,
	}
	winComputer2.print()
}
