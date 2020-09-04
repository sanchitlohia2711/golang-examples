package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

func main() {

	emp := employee{Name: "Sam", Age: 31, Salary: 2000}

	//Converting to jsonn
	empJSON, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(string(empJSON))
}
