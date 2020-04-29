package main

import "fmt"

func main() {
	lowercase := "abcdefghijklmnopqrstunwxyz"

	for _, c := range lowercase {
		fmt.Println(c)
	}

	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, c := range uppercase {
		fmt.Println(c)
	}

	numbers := "0123456789"
	for _, n := range numbers {
		fmt.Println(n)
	}
}
