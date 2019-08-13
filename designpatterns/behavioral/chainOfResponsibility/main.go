package main

func main() {

	patient := &patient{}

	cashier := &cashier{}
	medical := &medical{
		nextDepartment: cashier,
	}
	docter := &docter{
		nextDepartment: medical,
	}

	reception := &reception{
		nextDepartment: docter,
	}

	reception.execute(patient)
}
