package main

import (
	"fmt"
	"reflect"
)

func main() {
	var test interface{}
	test = "test_string"
	fmt.Println(reflect.TypeOf(test))
}
