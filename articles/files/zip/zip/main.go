package main

import (
	"archive/zip"
	"log"
	"os"
)

func main() {
	filesNamesToZip := []string{"a.txt", "b.txt"}
	var filesToZip []*os.File
	//Create files and write some content in it
	for _, fileName := range filesNamesToZip {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		filesToZip = append(filesToZip, file)
		file.WriteString("Writing some content to file")
	}
	//Create a zip file
	zipFile, err := os.Create("zipFile.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()

	zipW := zip.NewWriter(zipFile)
	defer zipW.Close()

	for _, file := range filesToZip {
		zipW.CreateHeader(file.H)
	}

}
