package main

import "fmt"

func main() {
	//Declare
	var employeeSalary map[string]int
	fmt.Println(employeeSalary)

	//Intialize using make
	employeeSalary2 := make(map[string]int)
	fmt.Println(employeeSalary2)

	//Intialize using map lieteral
	employeeSalary3 := map[string]int{
		"John": 1000,
		"Sam":  1200,
	}
	fmt.Println(employeeSalary3)

	//Operations
	//Add
	employeeSalary3["Carl"] = 1500

	//Get
	fmt.Printf("John salary is %d\n", employeeSalary3["John"])

	//Delete
	delete(employeeSalary3, "Carl")

	//Print map
	fmt.Println("\nPrinting employeeSalary3 map")
	fmt.Println(employeeSalary3)

}
