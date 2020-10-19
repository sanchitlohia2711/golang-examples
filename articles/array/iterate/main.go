package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c", "d", "e"}

	//With index and value
	for i, letter := range letters {
		fmt.Printf("%d %s", i, letter)
		fmt.Println()
	}

	//Only value
	for _, letter := range letters {
		fmt.Println(letter)
	}

	//Only index
	for i := range letters {
		fmt.Println(i)
	}

	//Without range and value. Just print array values
	i := 0
	for range letters {
		fmt.Println(i)
		i++
	}

	//Using conditional with for
	len := len(letters)
	for i := 0; i < len; i++ {
		fmt.Println(letters[i])
	}

	//Multiple initilization and post statement
	for i, j := 0, len; i < j; i, j = i+1, j-1 {
		fmt.Println(letters[i])
	}
}

func t(s []string) {
	gg := len(s)
}
