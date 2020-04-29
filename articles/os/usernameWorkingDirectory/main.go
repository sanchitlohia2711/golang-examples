package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	homeDirectory := user.HomeDir

	fmt.Printf("Home Directory: %s\n", homeDirectory)
}
