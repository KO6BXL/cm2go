package block

// If all inputs are off, output is on
func NOR() *Base {
	block := new(Base)
	block.name = "NOR"
	block.id = 0
	block.color = Color{R: 255, G: 9, B: 0}

	return block
}

// If all inputs are on, output in on
func AND() *Base {
	block := new(Base)
	block.name = "AND"
	block.id = 1
	block.color = Color{R: 0, G: 121, B: 255}

	return block
}

// If any inputs is on, output is on
func OR() *Base {
	block := new(Base)
	block.name = "OR"
	block.id = 2
	block.color = Color{R: 0, G: 241, B: 29}

	return block
}

// If an odd amount of inputs are on, output is on
func XOR() *Base {
	block := new(Base)
	block.name = "XOR"
	block.id = 3
	block.color = Color{R: 168, G: 0, B: 255}

	return block
}

// A button, output only
func BUTTON() *Base {
	block := new(Base)
	block.name = "BUTTON"
	block.id = 4
	block.color = Color{R: 255, G: 127, B: 0}

	return block
}

// If any input is on, output toggles between on and off
func FLIPFLOP() *Base {
	block := new(Base)
	block.name = "FLIPFLOP"
	block.id = 5
	block.color = Color{R: 30, G: 30, B: 30}

	return block
}

// Default: (175, 175, 175), 100, 25, false
type LEDParams struct {
	Color      Color
	OpacityOn  uint8
	OpacityOff uint8
	Analog     bool
}

// If any input is on, output is on, produces light. If analog is true, less inputs on, lower the opacity is
func LED(params LEDParams) *Base {
	block := new(Base)
	block.name = "LED"
	block.id = 6
	block.color = params.Color

	var analogOn int32 = 0
	if params.Analog {
		analogOn = 1
	}

	block.properties = []int32{
		int32(params.Color.R),
		int32(params.Color.G),
		int32(params.Color.B),
		int32(params.OpacityOn),
		int32(params.OpacityOff),
		analogOn,
	}

	return block
}

// If any input is on, output is on, produces sound. Intrument is: sine, square, triangle, sawtooth
func SOUND(frequency uint16, instrument uint8) *Base {
	block := new(Base)
	block.name = "SOUND"
	block.id = 7
	block.color = Color{R: 175, G: 131, B: 76}
	block.properties = []int32{int32(frequency), int32(instrument)}

	return block
}

// If any input is on, output is on, other conductors adjacent to it are also on
func CONDUCTOR() *Base {
	block := new(Base)
	block.name = "CONDUCTOR"
	block.id = 8
	block.color = Color{R: 73, G: 185, B: 255}

	return block
}

// If any input is on, output is on. Not intended for use
func CUSTOM() *Base {
	block := new(Base)
	block.name = "CUSTOM"
	block.id = 9
	block.color = Color{R: 255, G: 72, B: 72}

	return block
}

// If all inputs are off, output is on
func NAND() *Base {
	block := new(Base)
	block.name = "NAND"
	block.id = 10
	block.color = Color{R: 0, G: 42, B: 89}

	return block
}

// If an even number of inputs are on, output is on
func XNOR() *Base {
	block := new(Base)
	block.name = "XNOR"
	block.id = 11
	block.color = Color{R: 213, G: 0, B: 103}

	return block
}

// Produces random output each tick, output only
func RANDOM() *Base {
	block := new(Base)
	block.name = "RANDOM"
	block.id = 12
	block.color = Color{R: 84, G: 54, B: 35}

	return block
}

// If any input is on, output is on, intended for text
func TEXT(character uint8) *Base {
	block := new(Base)
	block.name = "TEXT"
	block.id = 13
	block.color = Color{R: 25, G: 71, B: 84}
	block.properties = []int32{int32(character)}

	return block
}

// If any input is on, output is on, intended for decoration
func TILE(color Color, material uint8) *Base {
	block := new(Base)
	block.name = "TILE"
	block.id = 14
	block.color = color
	block.properties = []int32{int32(color.R), int32(color.G), int32(color.B), int32(material)}

	return block
}

// If any input is on, output is on, instant
func NODE() *Base {
	block := new(Base)
	block.name = "NODE"
	block.id = 14
	block.color = Color{R: 165, G: 177, B: 200}

	return block
}

// If any input is on, output is on, has delay in ticks up to 1000
func DELAY(ticks uint16) *Base {
	block := new(Base)
	block.name = "DELAY"
	block.id = 14
	block.color = Color{R: 98, G: 24, B: 148}
	block.properties = []int32{int32(ticks)}

	return block
}

// If any input is on, channel is on. If channel is on, output is on
func ANTENNA(channel uint16) *Base {
	block := new(Base)
	block.name = "ANTENNA"
	block.id = 14
	block.color = Color{R: 235, G: 233, B: 183}
	block.properties = []int32{int32(channel)}

	return block
}

// If any input is on, output is on, other conductors adjacent to it are also on, instant
func CONDUCTORV2() *Base {
	block := new(Base)
	block.name = "CONDUCTORV2"
	block.id = 14
	block.color = Color{R: 52, G: 132, B: 182}

	return block
}
