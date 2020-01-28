package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	createSampleFiles()
	currentDirecoty, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	iterate(currentDirecoty)
}

func createSampleFiles() {
	fileNames := []string{"a.txt", "b.txt"}
	//Create files and write some content in it
	for _, fileName := range fileNames {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		file.WriteString("Writing some content to file")
	}
}

func 

func iterate(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("File Name: %s\n", info.Name())
		return nil
	})
}
