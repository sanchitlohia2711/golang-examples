package main

type director struct {
	builder *builder
}

func newDirector(b *builder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) buildHouse() house {
	d.builder.setDoorType("someDoorType")
	d.builder.setWindowType("someWindowType")
	d.builder.setNumFloor(2)
	return d.builder.getHouse()
}
