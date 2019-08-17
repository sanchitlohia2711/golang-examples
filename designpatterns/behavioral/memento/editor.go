package main

type editor struct {
	state string
}

func (e *editor) createMemento() *memento {
	return &memento{state: e.state}
}

func (e *editor) restoreState(m *memento) {
	e.state = m.getSavedState()
}

func (e *editor) setState(state string) {
	e.state = state
}

func (e *editor) getState() string {
	return e.state
}
