package main

import "sync"

type stationManager struct {
	isPlatformFree bool
	lock           *sync.Mutex
}

func newStationManger() *stationManager {
	return &stationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *stationManager) canLand() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	return false
}

func (s *stationManager) notifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
}
