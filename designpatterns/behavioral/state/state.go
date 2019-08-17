package main

type state interface {
	addItem() error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}
