package main

import "fmt"

func main() {
	//Declare
	employeeSalary := make(map[string]int)

	//Adding a key value
	employeeSalary["Tom"] = 2000

	fmt.Println("Key exists case")
	val, ok := employeeSalary["Tom"]
	fmt.Printf("Val: %d, ok: %t\n", val, ok)

	fmt.Println("Key doesn't exists case")
	val, ok = employeeSalary["Sam"]
	fmt.Printf("Val: %d, ok: %t\n", val, ok)
}
