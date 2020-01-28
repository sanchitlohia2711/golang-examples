package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	info, err := os.Stat("temp")
	if os.IsNotExist(err) {
		log.Fatal("File does not exist.")
	}
	if info.IsDir() {
		fmt.Println("temp is a directory")
	} else {
		fmt.Println("temp is a file")
	}
}
