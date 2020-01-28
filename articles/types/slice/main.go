package main

import "fmt"

func main() {
	//Declare a slice using make
	s := make([]string, 2, 3)
	fmt.Println(s)

	//Direct intialization
	p := []string{"a", "b", "c"}
	fmt.Println(p)

	//Append function
	p = append(p, "d")
	p = append(p, "d")
	p = append(p, "d")
	p = append(p, "d")
	fmt.Println(p)

	//Iterate over a slcie
	for _, val := range p {
		fmt.Println(val)
	}

}
