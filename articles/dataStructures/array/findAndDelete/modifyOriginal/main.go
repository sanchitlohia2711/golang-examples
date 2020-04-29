package main

import "fmt"

func main() {
	s := [4]int{1, 2, 3, 1}
	after := findAndDelete(s, 1)
	fmt.Println(after)

}

func findAndDelete(s [4]int, item int) []int {
	index := 0
	for _, i := range s {
		if i != item {
			s[index] = i
			index++
		}
	}
	return s[:index]
}
