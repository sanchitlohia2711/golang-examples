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
	folderInfo, err := os.Stat("temp")
	if os.IsNotExist(err) {
		log.Fatal("Folder does not exist.")
	}
	log.Println(folderInfo)
}
