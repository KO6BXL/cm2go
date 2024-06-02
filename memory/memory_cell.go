package memory

import "github.com/nameless9000/cm2go/block"

type MemoryCell struct {
	Output   *block.Base
	Input    *block.Base
	WriteBit *block.Base
}

// A basic memory cell using 1 FLIPFLOP, 1 AND, and 1 XOR.
// The XOR is the input and the AND is the write bit.
// Size: x=0-2
func NewMemoryCell() (block.Collection, MemoryCell) {
	var cell block.Collection

	flipflop := cell.Append(block.FLIPFLOP())
	and := cell.Append(block.AND())
	xor := cell.Append(block.XOR())

	and.Offset.X = 1
	xor.Offset.X = 2

	cell.Connect(xor, and)
	cell.Connect(and, flipflop)
	cell.Connect(flipflop, xor)

	return cell, MemoryCell{Output: flipflop, Input: xor, WriteBit: and}
}

type Register struct {
	Outputs  []*block.Base
	Inputs   []*block.Base
	WriteBit *block.Base
}

// Size: x=0-3 y=0-bits
func NewRegister(bits uint16) (collection block.Collection, register Register) {
	writeBit := collection.Append(block.NODE())
	writeBit.Offset.X = 3
	register.WriteBit = writeBit

	var bit uint16
	for ; bit < bits; bit++ {
		flipflop := collection.Append(block.FLIPFLOP())
		and := collection.Append(block.AND())
		xor := collection.Append(block.XOR())

		flipflop.Offset.Y = float32(bit)
		and.Offset.X = 1
		and.Offset.Y = float32(bit)
		xor.Offset.X = 2
		xor.Offset.Y = float32(bit)

		collection.Connect(xor, and)
		collection.Connect(and, flipflop)
		collection.Connect(flipflop, xor)

		register.Inputs = append(register.Inputs, xor)
		register.Outputs = append(register.Outputs, flipflop)
		collection.Connect(writeBit, and)
	}

	return
}
