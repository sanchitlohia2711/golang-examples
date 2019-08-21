package main

type visitor interface {
	visitAlpha(element)
	visitBeta(element)
}
