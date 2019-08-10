package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
)

//Test struct
type Test struct {
	Name string
	Age  int
}

func main() {
	test := Test{"Peter", 20}
	//fmt.Printf uses reflection to print test struct. We don't have to write any special code for the test type
	fmt.Printf("%v\n", test)
	fmt.Printf("%+v\n", test)
	fmt.Printf("%#v\n", test)

	//json.Marshal uses reflection to do valid encoding into json format. We don't have to write any special code for the test type
	data, err := json.Marshal(test)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)

	//json.Marshal uses reflection to do valid encoding into json format. We don't have to write any special code for the test type
	data, err = xml.MarshalIndent(test, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)

	//Here is a little glimpse of how above packages are using reflection to their benefits
	customPrint("Peter Age is ", 20, "\n")
}

func customPrint(args ...interface{}) {
	for _, arg := range args {
		switch v := reflect.ValueOf(arg); v.Kind() {
		case reflect.String:
			os.Stdout.WriteString(v.String())
		case reflect.Int:
			os.Stdout.WriteString(strconv.FormatInt(v.Int(), 10))
		}
	}
}
