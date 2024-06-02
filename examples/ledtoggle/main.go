package main

import (
	"github.com/nameless9000/cm2go/block"
	"github.com/nameless9000/cm2go/build"
)

func main() {
	var connections block.Connections
	var collection block.Collection

	defaultLedParams := block.LEDParams{Color: block.Color{R: 175, G: 175, B: 175}, OpacityOn: 100, OpacityOff: 25, Analog: false}

	toggle := collection.Append(block.FLIPFLOP())
	toggle.Offset.X = 1

	led := collection.Append(block.LED(defaultLedParams))

	connections.Connect(toggle, led)

	output, err := build.FastCompile([]block.Collection{collection}, []block.Connections{connections})
	if err != nil {
		panic(err)
	}

	println(output)
}
