package main

import (
	"fmt"
	"image"
	"log"
	"os"
)

func main() {
	imagePath := "/Users/slohia/Desktop/Sundayafternooncopy/test.jpg"
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	config, name, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Height %d\n", config.Height)
	fmt.Printf("Width %d\n", config.Width)
	fmt.Println(name)
}
