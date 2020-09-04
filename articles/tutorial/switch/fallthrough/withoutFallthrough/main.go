package main

import "fmt"

func main() {
	i := 9
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
	case i < 50:
		fmt.Println("i is less than 50")
	case i < 100:
		fmt.Println("i is less than 100")
	}
}
