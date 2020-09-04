package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name   string
	Age    int
	salary int
}

func main() {

	emp := employee{Name: "Sam", Age: 31, salary: 2000}

	//Marshal
	empJSON, err := json.Marshal(emp)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Marshal funnction output %s\n", string(empJSON))
	//Converting to jsonn
	empJSON, err = json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("MarshalIndent funnction output %s\n", string(empJSON))
}
