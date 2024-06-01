package block

type Offset struct {
	X float32
	Y float32
	Z float32
}

type Color struct {
	R uint8
	G uint8
	B uint8
}

type Base struct {
	name       string
	id         uint8
	properties []int32
	color      Color
	Offset     Offset
	State      bool
}

func (block Base) Name() string {
	return block.name
}

func (block Base) Id() uint8 {
	return block.id
}

func (block Base) Properties() []int32 {
	return block.properties
}

func (block Base) Color() Color {
	return block.color
}
