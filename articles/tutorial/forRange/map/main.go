package main

import "fmt"

func main() {
	sample := map[string]string{
		"a": "x",
		"b": "y",
	}

	//Iterating over all keys and values
	fmt.Println("Both Key and Value")
	for k, v := range sample {
		fmt.Printf("key :%s value: %s\n", k, v)
	}

	//Iterating over only keys
	fmt.Println("\nOnly keys")
	for k := range sample {
		fmt.Printf("key :%s\n", k)
	}

	//Iterating over only values
	fmt.Println("\nOnly values")
	for _, v := range sample {
		fmt.Printf("value :%s\n", v)
	}
}
