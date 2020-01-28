package main

import "fmt"

func main() {
	s := "0b£"

	//Convert string to run
	runeS := []rune(s)

	//This will print the Unicode Points
	fmt.Printf("%U\n", runeS)

	//This will the decimal value of Unicode Code Point
	fmt.Println(runeS)
}
