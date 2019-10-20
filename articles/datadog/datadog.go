package main

import (
	"log"
	"os"

	"github.com/DataDog/datadog-go/statsd"
)

// Client -> Stuct for client connection
type Client struct {
	Monitor *statsd.Client
}

var cl = &Client{}

//init datadog
func init() {
	c, err := statsd.NewBuffered(, 2)
	if err != nil {
		panic(err.Error())
		return
	}
	c.Namespace = "test-service"
	cl = &Client{
		Monitor: c,
	}
	return
}

func GetClient() *Client {
	return cl
}

//SendGauge -> Sends out metrics to datadog
func SendGauge(name string, value float64, tags []string, rate float64) bool {
	err := cl.Monitor.Gauge(name, value, tags, rate)
	if err != nil {
		log.Println("Cannot fire event to statsd.")
		log.Println(err)
		return false
	}
	return true
}

//SendGaugeWrapper wrapper over sending gauge events
func SendGaugeWrapper(name string, tags []string) error {
	goEnvironment, ok := os.LookupEnv("GOENV")
	if !ok {
		goEnvironment = "development"
	}

	tags = append(tags, "env:"+goEnvironment)

	err := cl.Monitor.Count(name, 1, tags, 1)
	if err != nil {
		return err
	}

	return nil
}
