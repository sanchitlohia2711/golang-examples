package main

import "fmt"

func main() {
	before := [4]int{1, 2, 3, 1}
	after := findAndDelete(before, 1)
	fmt.Println(after)

}

func findAndDelete(s [4]int, itemToDelete int) []int {
	var new [4]int
	index := 0
	for _, i := range s {
		if i != itemToDelete {
			new[index] = i
			index++
		}
	}
	return new[:index]
}
