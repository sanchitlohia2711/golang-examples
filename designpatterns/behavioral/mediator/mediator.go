package main

type mediator interface {
	canLand() bool
	notifyFree()
}
