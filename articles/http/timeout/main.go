package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	dialTimeout := time.Second * 5
	tLSHandshakeTimeout := time.Second * 5
	responseHeaderTimeout := time.Second * 5
	expectContinueTimeout := time.Second * 5
	idleConnectionTimeout := time.Second * 5
	clientTimeout := time.Nanosecond * 1

	dialContext := (&net.Dialer{
		Timeout: dialTimeout,
	}).DialContext

	tp := http.Transport{
		DialContext:           dialContext,
		TLSHandshakeTimeout:   tLSHandshakeTimeout,
		ResponseHeaderTimeout: responseHeaderTimeout,
		ExpectContinueTimeout: expectContinueTimeout,
		IdleConnTimeout:       idleConnectionTimeout,
	}

	client := http.Client{
		Transport: &tp,
		Timeout:   clientTimeout,
	}

	req, err := http.NewRequest("GET", "https://www.golangbyexample.com", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(body))
}
