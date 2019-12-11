package main

import (
	"container/list"
	"fmt"
)

type customList struct {
	container *list.List
}

func (c *customList) AddFront(value string) {
	c.container.PushFront(value)
}

func (c *customList) RemoveFront() error {
	if c.container.Len() > 0 {
		ele := c.container.Front()
		c.container.Remove(ele)
	}
	return fmt.Errorf("Remove Front Error: Linked List is empty")
}

func (c *customList) Front() (string, error) {
	if c.container.Len() > 0 {
		if val, ok := c.container.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("Peep Error: Queue Datatype is incorrect")
	}
	return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customList) Size() int {
	return c.container.Len()
}

func (c *customList) Empty() bool {
	return c.container.Len() == 0
}

func main() {
	customQueue := &customList{
		container: list.New(),
	}

	fmt.Printf("Enqueue: A\n")
	customQueue.AddFront("A")
	fmt.Printf("Enqueue: B\n")
	customQueue.AddFront("B")
	fmt.Printf("Size: %d\n", customQueue.Size())

	for customQueue.Size() > 0 {
		frontVal, _ := customQueue.Front()
		fmt.Printf("Front: %s\n", frontVal)
		fmt.Printf("Dequeue: %s\n", frontVal)
		customQueue.RemoveFront()
	}
	fmt.Printf("Size: %d\n", customQueue.Size())
}
