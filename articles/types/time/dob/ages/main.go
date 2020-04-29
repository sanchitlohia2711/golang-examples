package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	oldTime := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	diff := now.Sub(oldTime)

	//In hours
	fmt.Printf("Hours: %f\n", diff.Hours())

	//In minutes
	fmt.Printf("Minutes: %f\n", diff.Minutes())

	//In seconds
	fmt.Printf("Seconds: %f\n", diff.Seconds())

	//In nanoseconds
	fmt.Printf("Nanoseconds: %d\n", diff.Nanoseconds())

	//time Since
	fmt.Println(time.Since(now.Add(-time.Hour * 1)))

	//time until
	fmt.Println(time.Until(now.Add(time.Hour * 1)))
}
