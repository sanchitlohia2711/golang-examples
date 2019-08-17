package main

import "fmt"

func main() {

	caretaker := &caretaker{
		mementoArray: make([]*memento, 0),
	}

	editor := &editor{
		state: "A",
	}

	
	fmt.Printf("Editor Current State: %s\n", editor.getState())
	caretaker.addMemento(editor.createMemento())

	editor.setState("B")
	fmt.Printf("Editor Current State: %s\n", editor.getState())
	caretaker.addMemento(editor.createMemento())

	editor.setState("C")
	fmt.Printf("Editor Current State: %s\n", editor.getState())
	caretaker.addMemento(editor.createMemento())

	editor.restoreState(caretaker.getMenento(1))
	fmt.Printf("Restored to State: %s\n", editor.getState())

	editor.restoreState(caretaker.getMenento(0))
	fmt.Printf("Restored to State: %s\n", editor.getState())

}
