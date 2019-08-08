package main

type builder struct {
	windowType string
	doorType   string
	floor      int
}

func newBuilder() *builder {
	return &builder{}
}

func (b *builder) setWindowType(w string) {
	b.windowType = w
}

func (b *builder) setDoorType(d string) {
	b.doorType = d
}

func (b *builder) setNumFloor(n int) {
	b.floor = n
}

func (b *builder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
