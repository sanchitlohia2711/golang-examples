package main

import (
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	configureHystrix()
	for i := 0; i < 10; i++ {
		go do(i)
	}
	fmt.Scanln()
}

func do(inex int) {
	time.Sleep(time.Second * time.Duration(inex))
	hystrix.Do("my_command", func() error {
		// talk to other services
		if inex > 5 {
			fmt.Printf("success + %d", inex)
			fmt.Println("")
			return nil
		}
		return fmt.Errorf("some + %d", inex)
	}, func(err error) error {
		// do this when services are down
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Success")
		}
		return nil
	})
}

func configureHystrix() {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:                1,
		MaxConcurrentRequests:  20,
		ErrorPercentThreshold:  90,
		RequestVolumeThreshold: 1,
		SleepWindow:            1000,
	})
}
