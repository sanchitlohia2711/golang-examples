package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	fmt.Println("herer")
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 100 * time.Second
	bodyGenerate := requestDTOV1.GenerateOTPParams{}
	bodyGenerate.Phone = "082213909101"
	bodyGenerate.Name = "Sanchit Lohia"
	bodyGenerate.Email = "sanchitloves1@gmail.com"
	bodyGenerate.CountryCode = "62"
	bodyGenerateBytes, _ := json.Marshal(bodyGenerate)

	headers := make(map[string][]string)
	headers["Content-Type"] = []string{"application/json"}
	headers["X-Client-Id"] = []string{"tokopedia"}

	body := requestDTOV1.VerifyOTPParams{}
	body.OTP = "9101"
	body.RefID = "1234-1234"
	body.Action = "otpLinkage"
	bodyBytes, _ := json.Marshal(body)

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "http://localhost:8080/v1/oauth/otp/generate?msgId=123",
		Body:   bodyGenerateBytes,
		Header: http.Header(headers),
	}, vegeta.Target{Method: "POST",
		URL:    "http://localhost:8080/v1/oauth/otp/verify?msgId=123",
		Body:   bodyBytes,
		Header: http.Header(headers)})
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
