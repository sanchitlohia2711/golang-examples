package main

import "fmt"

func main() {
	before := []int{1, 2, 3, 1}
	after := findAndDelete(before, 1)
	fmt.Println(after)

}

func findAndDelete(s []int, itemToDelete int) []int {
	var new = make([]int, len(s))
	index := 0
	for _, i := range s {
		if i != itemToDelete {
			new = append(new, i)
			index++
		}
	}
	return new[:index]
}
