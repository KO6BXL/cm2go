package main

import (
	"github.com/nameless9000/cm2go/block"
	"github.com/nameless9000/cm2go/build"
	"github.com/nameless9000/cm2go/memory"
)

func memoryCell() {
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

func register() {
	var bits uint16 = 8

	collection, register := memory.NewRegister(bits)
	collection.Position.Z = 1

	var inputCollection block.Collection

	writeBit := inputCollection.Append(block.BUTTON())
	writeBit.Offset.X = 3
	inputCollection.Connect(writeBit, register.WriteBit)

	for bit, input := range register.Inputs {
		flipflop := inputCollection.Append(block.FLIPFLOP())
		flipflop.Offset.X = 2
		flipflop.Offset.Y = float32(bit)

		inputCollection.Connect(flipflop, input)
	}

	for bit, output := range register.Outputs {
		led := inputCollection.Append(block.LED(nil))
		led.Offset.Y = float32(bit)

		inputCollection.Connect(output, led)
	}

	output, err := build.FastCompile([]block.Collection{collection, inputCollection})
	if err != nil {
		panic(err)
	}

	println(output)
}

func main() {
	println("Memory Cell:")
	memoryCell()
	println("\nRegister:")
	register()
}
