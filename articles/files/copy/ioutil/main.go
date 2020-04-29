package main

import (
	"io/ioutil"
	"log"
)

// Copy a file
func main() {
	//Read all the contents of the  original file
	bytesRead, err := ioutil.ReadFile("original.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Copy all the contents to the desitination file
	err = ioutil.WriteFile("new.txt", bytesRead, 0755)
	if err != nil {
		log.Fatal(err)
	}

}
