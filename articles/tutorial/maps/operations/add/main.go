package main

import "fmt"

func main() {
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary)
}
