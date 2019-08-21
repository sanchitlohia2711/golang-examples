package main

type element interface {
	visit(visitor)
}
