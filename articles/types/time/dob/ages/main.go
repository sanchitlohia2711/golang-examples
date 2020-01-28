package main

import (
	"fmt"
	"time"

	age "github.com/bearbin/go-age"
)

func main() {
	dob := getDOB(2011, 4, 2)
	fmt.Printf("Age is %d\n", age.Age(dob))
}

func getDOB(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}
