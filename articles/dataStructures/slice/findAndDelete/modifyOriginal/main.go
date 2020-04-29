package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 1}
	after := findAndDelete(s, 1)
	fmt.Println(after)

}

func findAndDelete(s []int, item int) []int {
	index := 0
	for _, i := range s {
		if i != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}
