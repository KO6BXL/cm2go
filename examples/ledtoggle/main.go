package main

import (
	"github.com/nameless9000/cm2go/block"
	"github.com/nameless9000/cm2go/build"
)

func main() {
	var connections block.Connections
	var collection block.Collection

	toggle := collection.Append(block.FLIPFLOP())
	toggle.Offset.X = 1

	led := collection.Append(block.LED(nil))

	connections.Connect(toggle, led)

	output, err := build.FastCompile([]block.Collection{collection}, []block.Connections{connections})
	if err != nil {
		panic(err)
	}

	println(output)
}
