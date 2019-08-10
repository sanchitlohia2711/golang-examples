package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func startLoadTest() {
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 1 * time.Second

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8080/app/status",
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("Earliest: %v\n", metrics.Earliest)
	fmt.Printf("Latest: %v\n", metrics.Latest)
	fmt.Printf("End: %v\n", metrics.End)
	fmt.Printf("Duration: %d\n", metrics.Duration)
	fmt.Printf("Wait: %d\n", metrics.Wait)
	fmt.Printf("Requests: %d\n", metrics.Requests)
	fmt.Printf("Rate: %f\n", metrics.Rate)
	fmt.Printf("StatusCodes: +%v\n", metrics.StatusCodes)
	fmt.Printf("Success: %f\n", metrics.Success)
	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	fmt.Printf("95th percentile: %s\n", metrics.Latencies.P95)
	fmt.Printf("50th percentile: %s\n", metrics.Latencies.P50)
	fmt.Printf("Mean: %s\n", metrics.Latencies.Mean)
	fmt.Printf("Max: %s\n", metrics.Latencies.Max)
}
