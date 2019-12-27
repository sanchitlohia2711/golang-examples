package main

import (
	"context"
	"fmt"
	"time"
)

type shopkeeper struct {
	customer *customer
	bank     *bank
}

func (s *shopkeeper) makePayment(ctx context.Context, cardDetails string, price int) {
	paymentDoneChannel := s.bank.initPayment(cardDetails, price)
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err())
			return
		case <-paymentDoneChannel:
			fmt.Println("Payment Completed exit")
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}

type bank struct {
	shopkeeper shopkeeper
}

func (b *bank) initPayment(cardDetails string, price int) chan struct{} {
	done := make(chan struct{}, 1)
	go b.receievePayment(done, cardDetails, price)
	return done
}

func (b *bank) receievePayment(done chan struct{}, cardDetails string, price int) {
	time.Sleep(1 * time.Second)
	done <- struct{}{}
	return
}

type customer struct {
	name        string
	price       int
	cardDetails string
}

func main() {
	ctx := context.Background()

	cancelCtx, cancelFunc := context.WithCancel(ctx)

	customer := &customer{
		name:        "test",
		cardDetails: "1234",
	}

	bank := &bank{}

	shopkeeper := shopkeeper{
		customer: customer,
		bank:     bank,
	}

	go shopkeeper.makePayment(cancelCtx, customer.cardDetails, 10)

	time.Sleep(time.Second * 3)
	cancelFunc()
	time.Sleep(time.Second * 1)
}

func makePayment(ctx context.Context) {
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
