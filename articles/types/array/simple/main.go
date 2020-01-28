package main

import "fmt"

func main() {
	//Declare a array
	sample := [3]string{"a", "b", "c"}
	print(sample)

}

func print(sample [3]string) {
	fmt.Println(sample)
}
