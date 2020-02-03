package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC)
	fmt.Println(t)
	time.Now()
}
