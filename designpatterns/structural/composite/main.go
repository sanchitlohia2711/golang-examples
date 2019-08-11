package main

import "fmt"

func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		compositeArray: []composite{file1},
		name:           "Folder1",
	}
	fmt.Println("Printing hierarchy for Folder1")
	folder1.print("  ")

	folder2 := &folder{
		compositeArray: []composite{folder1, file2, file3},
		name:           "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")
}
