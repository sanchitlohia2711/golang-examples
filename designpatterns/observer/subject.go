package main

type subject interface {
	registerObserver(Observer observer)
	deregisterObserver(Observer observer)
	notifyObserver()
}
