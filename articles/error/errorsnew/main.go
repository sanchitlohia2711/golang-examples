package main

import (
	"errors"
	"fmt"
)

func main() {
	sampleErr := errors.New("error occured")

	fmt.Println(sampleErr)
}
