package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

//Reader read file using bufio reader
func Reader() {
	file, err := os.Open("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return
		}
		_ = line // GET the line string
	}
}
