package main

import (
	"fmt"
)

func main() {
	sampleErr := fmt.Errorf("Err is: %s", "database connection issue")

	fmt.Println(sampleErr)
}
