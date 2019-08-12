package main

import "fmt"

func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristRedDressType)
	game.addTerrorist(TerroristBlackDressType)
	game.addTerrorist(TerroristRedDressType)
	game.addTerrorist(TerroristRedDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerrroristGrayDressType)
	game.addCounterTerrorist(CounterTerrroristGreenDressType)
	game.addCounterTerrorist(CounterTerrroristGrayDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
