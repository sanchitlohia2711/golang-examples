package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	colorRed := "\033[31m"
	colorReset := "\033[0m"
	color.Cyan("print C")

	fmt.Println(string(colorRed), "test", string(colorReset))
	fmt.Println("next")
}
