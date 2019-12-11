package main

import "fmt"

type ele struct {
	name string
	next *ele
}

type singleList struct {
	len  int
	head *ele
}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) addFront(name string) {
	ele := &ele{
		name: name,
	}
	if s.head == nil {
		s.head = ele
	} else {
		ele.next = s.head
		s.head = ele
	}
	s.len++
	return
}

func (s *singleList) addBack(name string) {
	ele := &ele{
		name: name,
	}
	if s.head == nil {
		s.head = ele
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = ele
	}
	s.len++
	return
}

func (s *singleList) front() (string, error) {
	if s.head == nil {
		return "", fmt.Errorf("Single List is empty")
	}
	return s.head.name, nil
}

func (s *singleList) removeFront() error {
	if s.head == nil {
		return fmt.Errorf("List is empty")
	}
	s.head = s.head.next
	s.len--
	return nil
}

func (s *singleList) removeBack() error {
	if s.head == nil {
		return fmt.Errorf("RemoveBack: List is empty")
	}
	var prev *ele
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		s.head = nil
	}
	s.len--
	return nil
}

func (s *singleList) size() int {
	return s.len
}

func (s *singleList) traverse() error {
	if s.head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}
	return nil
}

func main() {
	singleList := initList()

	fmt.Printf("AddFront: A\n")
	singleList.addFront("A")
	fmt.Printf("AddFront: B\n")
	singleList.addFront("B")
	fmt.Printf("AddBack: C\n")
	singleList.addBack("C")

	err := singleList.traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("RemoveFront\n")
	err = singleList.removeFront()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.removeBack()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.removeBack()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.removeBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	err = singleList.traverse()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Size: %d\n", singleList.size())
}
