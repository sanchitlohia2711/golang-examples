package main

type observer interface {
	notify(newStockValue int)
	getEmail() string
}
