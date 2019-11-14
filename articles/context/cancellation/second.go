package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func mains() {
	ctx := context.Background()

	cancelCtx, cancelFunc := context.WithCancel(ctx)

	go task1(cancelCtx)

	err := task2(cancelCtx)
	if err != nil {
		cancelFunc()
	}
	time.Sleep(time.Second * 1)
}

func task2(ctx context.Context) error {
	time.Sleep(time.Second * 3)
	return errors.New("Error occured")
}

func ttask1(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
