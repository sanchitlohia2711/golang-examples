package main

import (
	"fmt"
	"sync"
)

type pool struct {
	idle     []poolObject
	active   []poolObject
	capacity int
	mulock   *sync.Mutex
}

//InitPool Initialize the pool
func InitPool(poolObjects []poolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("Cannot craete a pool of 0 length")
	}
	active := make([]poolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}

	return pool, nil
}

func (p *pool) loan() (poolObject, error) {
	p.mulock.Lock()

	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No pool object free. Please request after sometime")
	}

	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	p.mulock.Unlock()

	return obj, nil
}

func (p *pool) receive(target poolObject) error {
	p.mulock.Lock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.active = append(p.active, target)
	p.mulock.Unlock()
	return nil
}

func (p *pool) remove(target poolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getId() == target.getId() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("Targe pool object doesn't belong to the pool")
}
