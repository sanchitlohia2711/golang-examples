package main

type vendingMachine struct {
	initState     state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount int
	itemPrice int
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem()
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}
