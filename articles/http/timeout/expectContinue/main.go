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
	net.Dial()
	dialTimeout := time.Second * 5
	tLSHandshakeTimeout := time.Second * 5
	responseHeaderTimeout := time.Second * 5
	expectContinueTimeout := time.Nanosecond * 1
	idleConnectionTimeout := time.Second * 5
	clientTimeout := time.Second * 5

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

	req, err := http.NewRequest("POST", "https://www.facebook.com/sanchit.lohia", nil)
	req.Header.Add("Expect", "100-Continue")
	if err != nil {
		log.Fatalf(err.Error())
	}
	net.Listen
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
