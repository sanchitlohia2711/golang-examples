package main

import (
	"fmt"
	"reflect"
)

func main() {
	sample := map[string]string{
		"a": "x",
		"b": "y",
	}

	//Iterating over all keys and values
	for k, v := range sample {
		fmt.Printf("key :%s value: %s\n", k, v)
	}

	//Iterating over only the keys
	for k := range sample {
		fmt.Printf("key :%s\n", k)
	}

	//Iterating over only values
	for _, v := range sample {
		fmt.Printf("value :%s\n", v)
	}

	//Get list of all keys
	keys := getAllKeys(sample)
	fmt.Println(keys)

	s := reflect.ValueOf(sample).MapKeys()
	fmt.Println(s)

}

func getAllKeys(sample map[string]string) []string {
	var keys []string
	for k := range sample {
		keys = append(keys, k)
	}
	return keys
}
