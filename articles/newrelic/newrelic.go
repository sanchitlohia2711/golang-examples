package main

import (
	newrelic "github.com/newrelic/go-agent"
)

var (
	nrapp newrelic.Application
)

func initNewRelic() {
	var err error
	nrConfig := newrelic.NewConfig("test", "somekey")
	nrapp, err = newrelic.NewApplication(nrConfig)
	if err != nil {
		panic("Failed to setup NewRelic: " + err.Error())
	}
}
