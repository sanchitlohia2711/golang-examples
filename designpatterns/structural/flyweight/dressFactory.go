package main

import "fmt"

const (
	//TerroristRedDressType terrorist red dress type
	TerroristRedDressType = "tRedDress"
	//TerroristBlackDressType terrorist black dress type
	TerroristBlackDressType = "tBlackDress"
	//CounterTerrroristGreenDressType terrorist green dress type
	CounterTerrroristGreenDressType = "ctGreenDress"
	//CounterTerrroristGrayDressType terrorist gray dress type
	CounterTerrroristGrayDressType = "ctGrayDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristRedDressType {
		d.dressMap[dressType] = &terroristDress{
			color: "red",
		}
		return d.dressMap[dressType], nil
	}
	if dressType == TerroristBlackDressType {
		d.dressMap[dressType] = &terroristDress{
			color: "black",
		}
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerrroristGreenDressType {
		d.dressMap[dressType] = &counterTerroristDress{
			color: "green",
		}
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerrroristGrayDressType {
		d.dressMap[dressType] = &counterTerroristDress{
			color: "gray",
		}
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}
