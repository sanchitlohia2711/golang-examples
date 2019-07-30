package main

type stockData struct {
	observerList []observer
	stockValue   int
}

func newStockData(stockValue int) *stockData {
	return &stockData{
		stockValue: stockValue,
	}
}

func (s *stockData) updateStockValue(newStockValue int) {
	s.stockValue = newStockValue
}
func (s *stockData) registerObserver(o observer) {
	s.observerList = append(s.observerList, o)
}

func (s *stockData) deregisterObserver(o observer) {
	s.observerList = removeFromslice(s.observerList, o)
}

func (s *stockData) notifyObserver() {
	for _, observer := range s.observerList {
		observer.notify(s.stockValue)
	}
}

func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getEmail() == observer.getEmail() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
