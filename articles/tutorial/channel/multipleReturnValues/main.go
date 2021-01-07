package main

import (
	"fmt"
	"time"
)

type result struct {
	sumValue      int
	multiplyValue int
}

func main() {
	resultChan := make(chan result, 1)
	sumAndMultiply(2, 3, resultChan)

	res := <-resultChan
	fmt.Printf("Sum Value: %d\n", res.sumValue)
	fmt.Printf("Multiply Value: %d\n", res.multiplyValue)
	close(resultChan)

}

func sumAndMultiply(a, b int, resultChan chan result) {
	sumValue := a + b
	multiplyValue := a * b
	res := result{sumValue: sumValue, multiplyValue: multiplyValue}
	resultChan <- res
	return
}
