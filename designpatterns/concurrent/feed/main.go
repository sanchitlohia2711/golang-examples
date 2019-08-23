package main

import (
	"fmt"
	"time"
)

type Item struct {
	Title, Channel, GUID string
}
type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

func Fetch(domain string) Fetcher {

}

type Subscription interface {
	Updates() <-chan Item
	Close() error
}

func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),
	}

	go s.loop()
	return s
}

type sub struct {
	fetcher Fetcher
	updates chan Item
}

func (s *sub) loop() {

}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	return nil
}
func Merge(subs ...Subscription) Subscription {

}

func main() {
	merged := Merge(
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("blog.golang.org")),
	)

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.Close())
	})

	for it := range merged.Updates() {
		fmt.Println(it.Channel, it.Title)
	}
}
