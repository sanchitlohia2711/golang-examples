package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	oldTime := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	diff := currentTime.Sub(oldTime)

	//In hours
	fmt.Printf("Hours: %f\n", diff.Hours())

	//In minutes
	fmt.Printf("Minutes: %f\n", diff.Minutes())

	//In seconds
	fmt.Printf("Seconds: %f\n", diff.Seconds())

	//In nanoseconds
	fmt.Printf("Nanoseconds: %d\n", diff.Nanoseconds())
}
