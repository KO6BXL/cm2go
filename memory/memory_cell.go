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
