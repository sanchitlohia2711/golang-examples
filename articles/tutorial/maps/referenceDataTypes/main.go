package main

import "fmt"

func main() {
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	employeeSalary["Tom"] = 2000
	employeeSalary["Sam"] = 1200

	eS := employeeSalary

	//Change employeeSalary
	employeeSalary["John"] = 3000
	fmt.Println("Changing employeeSalary Map")
	fmt.Printf("employeeSalary: %v\n", employeeSalary)
	fmt.Printf("eS: %v\n", eS)

	//Change eS
	employeeSalary["John"] = 4000
	fmt.Println("\nChanging eS Map")
	fmt.Printf("employeeSalary: %v\n", employeeSalary)
	fmt.Printf("eS: %v\n", eS)

}
