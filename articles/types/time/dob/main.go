package main

import (
	"fmt"
	"time"
)

const (
	//TimeFormat1 to format date into
	TimeFormat1 = "2006-01-02"
	//TimeFormat2 Other format to format date time
	TimeFormat2 = "January 02, 2006"
)

func main() {
	dob := getDOB(2011, 4, 2)
	fmt.Println(dob.Format(TimeFormat1))
	fmt.Println(dob.Format(TimeFormat2))
}

func getDOB(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}
