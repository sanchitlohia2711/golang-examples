package main

import (
	"fmt"
	"time"
)

func main() {
	//time.Now() example
	t := time.Now()
	fmt.Println(t)

	//time.Date() Example
	t = time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC)
	fmt.Println(t)

	//time.Parse() Example
	//Parse YYYY-MM-DD
	t, _ = time.Parse("2006-01-02", "2020-01-29")
	fmt.Println(t)
}
