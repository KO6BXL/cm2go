package main

import (
	"github.com/ko6bxl/cm2go/block"
	"github.com/ko6bxl/cm2go/build"
)

func main() {
	var collection block.Collection

	toggle := collection.Append(block.FLIPFLOP())
	toggle.Offset.X = 1

	led := collection.Append(block.LED(nil))

	collection.Connect(toggle, led)

	output, err := build.FastCompile([]block.Collection{collection})
	if err != nil {
		panic(err)
	}

	println(output)
}
