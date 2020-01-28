package main

import (
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	fileinfo, err := os.Stat("temp.txt")
	if os.IsNotExist(err) {
		log.Fatal("File does not exist.")
	}
	log.Println(fileinfo)
}
