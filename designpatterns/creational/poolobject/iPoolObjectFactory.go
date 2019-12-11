package main

type iPoolObjectFactory interface {
	create() (iPoolObject, error)
}
