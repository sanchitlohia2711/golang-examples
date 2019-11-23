package main

import (
	"fmt"
	"time"
)

func main() {
	logger := GetLogger()

	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 1)
		logger.Infof("test %d", i)
	}

	fmt.Scanln()
}
