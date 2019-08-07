package main

import "fmt"

//Observable struct
type Observable struct {
	email string
}

func (o *Observable) notify(stockValue int) {
	fmt.Printf("Stock value is %d", stockValue)
}

func (o *Observable) getEmail() string {
	return o.email
}

func main() {

	fmt.Println("df")
	p := 2
	stockDataSubject := newStockData(p)

	observerFirst := &Observable{email: "abc@gmail.com"}
	observerSecond := &Observable{email: "xyz@gmail.com"}

	stockDataSubject.registerObserver(observerFirst)
	stockDataSubject.registerObserver(observerSecond)

	stockDataSubject.updateStockValue(5)
}
