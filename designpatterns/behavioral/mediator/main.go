package main

func main() {

	stationManager := newStationManger()
	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}
	goodsTrain := &goodsTrain{
		mediator: stationManager,
	}

	passengerTrain.arrival()
	goodsTrain.arrival()

	passengerTrain.departure()
	goodsTrain.arrival()
}
