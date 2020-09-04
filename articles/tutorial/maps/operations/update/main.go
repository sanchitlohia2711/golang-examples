package main

import "fmt"

func main() {
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	fmt.Println("Before update")
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary)

	fmt.Println("After update")
	employeeSalary["Tom"] = 3000
	fmt.Println(employeeSalary)
}
