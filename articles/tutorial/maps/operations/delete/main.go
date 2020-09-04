package main

import "fmt"

func main() {
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	fmt.Println("Adding key")
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary)

	fmt.Println("\nDeleting key")
	delete(employeeSalary, "Tom")
	fmt.Println(employeeSalary)
}
