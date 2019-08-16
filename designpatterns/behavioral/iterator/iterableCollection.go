package main

type iterableCollection interface {
	createIterator() iterator
}
