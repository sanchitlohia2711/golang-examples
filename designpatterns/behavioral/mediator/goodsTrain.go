package main

import "fmt"

type goodsTrain struct {
	mediator mediator
}

func (g *goodsTrain) arrival() {
	if g.mediator.canLand() {
		fmt.Println("GoodsTrain: Landing")
	} else {
		fmt.Println("GoodsTrain: Waiting")
	}
}

func (g *goodsTrain) departure() {
	g.mediator.notifyFree()
	fmt.Println("GoodsTrain: Leaving")
}
