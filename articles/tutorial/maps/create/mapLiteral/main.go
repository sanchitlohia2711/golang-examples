package main

import "fmt"

func main() {
	//Declare
	employeeSalary := map[string]int{}
	fmt.Println(employeeSalary)

	//Intialize using map lieteral
	employeeSalary = map[string]int{
		"John": 1000,
		"Sam":  1200,
	}

	//Adding a key value
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary)
}
