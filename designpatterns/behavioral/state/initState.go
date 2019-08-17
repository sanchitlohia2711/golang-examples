package main

import "fmt"

type initState struct {
	vendingMachine *vendingMachine
}

func (i *initState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Init: item requestd")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *initState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.itemCount = i.vendingMachine.itemCount + count
	return nil
}

func (i *initState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}
func (i *initState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}
