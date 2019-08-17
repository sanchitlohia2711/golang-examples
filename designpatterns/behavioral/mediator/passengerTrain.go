package main

import "fmt"

type passengerTrain struct {
	mediator mediator
}

func (g *passengerTrain) arrival() {
	if g.mediator.canLand() {
		fmt.Println("PassengerTrain: Landing")
	} else {
		fmt.Println("PassengerTrain: Waiting")
	}
}

func (g *passengerTrain) departure() {
	g.mediator.notifyFree()
	fmt.Println("PassengerTrain: Leaving")
}
