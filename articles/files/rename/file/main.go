package main

import (
	"log"
	"os"
)

func main() {
	//Create a file
	file, err := os.Create("temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Change permission so that it can be moved
	err = os.Chmod("temp.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	err = os.Rename("temp.txt", "newTemp.txt")
	if err != nil {
		log.Fatal(err)
	}
}
