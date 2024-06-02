package main

import (
	"github.com/nameless9000/cm2go/block"
	"github.com/nameless9000/cm2go/build"
	"github.com/nameless9000/cm2go/memory"
)

func main() {
	cellCollection, cell := memory.NewMemoryCell()
	cellCollection.Position.Y = 3

	var inputCollection block.Collection

	led := inputCollection.Append(block.LED(nil))
	input := inputCollection.Append(block.FLIPFLOP())
	write := inputCollection.Append(block.FLIPFLOP())

	inputCollection.Connect(cell.Output, led)
	inputCollection.Connect(input, cell.Input)
	inputCollection.Connect(write, cell.WriteBit)

	write.Offset.X = 1
	input.Offset.X = 2

	output, err := build.FastCompile([]block.Collection{cellCollection, inputCollection})
	if err != nil {
		panic(err)
	}

	println(output)
}
