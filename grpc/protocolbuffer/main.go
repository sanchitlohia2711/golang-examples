package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	protoBufferExample()
	jsonExample()
	xmlExample()
}

func protoBufferExample() {
	person := &Person1{Name: "XXX"}
	fmt.Printf("Person1's name is %s\n", person.GetName())

	//Now lets write this person object to file
	t := time.Now()
	out, err := proto.Marshal(person)
	if err != nil {
		log.Fatalf("Serialization error: %s", err.Error())
	}
	fmt.Printf("ProtoBuffer Marshal Time: %d ns\n", time.Now().Sub(t).Nanoseconds())

	t = time.Now()
	if err := ioutil.WriteFile("person.bin", out, 0644); err != nil {
		log.Fatalf("Write File Error: %s ", err.Error())
	}
	fmt.Printf("ProtoBuffer Write File Time: %d ns\n", time.Now().Sub(t).Nanoseconds())

	//Read from file
	t = time.Now()
	in, err := ioutil.ReadFile("person.bin")
	if err != nil {
		log.Fatalf("Read File Error: %s ", err.Error())
	}
	fmt.Printf("ProtoBuffer Read File Time: %d ns\n", time.Now().Sub(t).Nanoseconds())
	person2 := &Person1{}
	t = time.Now()
	err2 := proto.Unmarshal(in, person2)
	if err2 != nil {
		log.Fatalf("DeSerialization error: %s", err.Error())
	}
	fmt.Printf("ProtoBuffer Unmarshal File Time: %d ns\n", time.Now().Sub(t).Nanoseconds())
}

func jsonExample() {
	person := Person1{Name: "XXX"}

	//Json Marshal
	t := time.Now()
	out, err := json.Marshal(person)
	if err != nil {
		log.Fatalf("Serialization error: %s", err.Error())
	}
	fmt.Printf("JSON Marshal Time: %d ns\n", time.Now().Sub(t).Nanoseconds())

	//Write to file
	t = time.Now()
	if err := ioutil.WriteFile("person.json", out, 0644); err != nil {
		log.Fatalf("Write File Error: %s ", err.Error())
	}
	fmt.Printf("JSON Write File Time: %d ns\n", time.Now().Sub(t).Nanoseconds())

	//Read from file
	t = time.Now()
	in, err := ioutil.ReadFile("person.json")
	if err != nil {
		log.Fatalf("Read File Error: %s ", err.Error())
	}
	fmt.Printf("JSON Read File Time: %d ns\n", time.Now().Sub(t).Nanoseconds())

	person2 := &Person1{}
	//Unmarshal
	t = time.Now()
	err = json.Unmarshal(in, person2)
	if err != nil {
		log.Fatalf("DeSerialization error: %s", err.Error())
	}
	fmt.Printf("JSON Unmarshal File Time: %d ns\n", time.Now().Sub(t).Nanoseconds())
}

func xmlExample() {
	person := Person1{Name: "XXX"}

	//Json Marshal
	t := time.Now()
	out, err := xml.Marshal(person)
	if err != nil {
		log.Fatalf("Serialization error: %s", err.Error())
	}
	fmt.Printf("XML Marshal Time: %d ns\n", time.Now().Sub(t).Nanoseconds())

	//Write to file
	t = time.Now()
	if err := ioutil.WriteFile("person.xml", out, 0644); err != nil {
		log.Fatalf("Write File Error: %s ", err.Error())
	}
	fmt.Printf("XML Write File Time: %d ns\n", time.Now().Sub(t).Nanoseconds())

	//Read from file
	t = time.Now()
	in, err := ioutil.ReadFile("person.xml")
	if err != nil {
		log.Fatalf("Read File Error: %s ", err.Error())
	}
	fmt.Printf("XML Read File Time: %d ns\n", time.Now().Sub(t).Nanoseconds())

	person2 := &Person1{}
	//Unmarshal
	t = time.Now()
	err = xml.Unmarshal(in, person2)
	if err != nil {
		log.Fatalf("DeSerialization error: %s", err.Error())
	}
	fmt.Printf("XML Unmarshal File Time: %d ns\n", time.Now().Sub(t).Nanoseconds())
}
