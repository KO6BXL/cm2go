package main

import (
	"fmt"

	"github.com/nameless9000/cm2go/block"
)

func main() {
	var connections block.Connections
	var collection block.Collection

	toggle := collection.Append(block.FLIPFLOP())
	led := collection.Append(block.LED(block.Color{R: 175, G: 175, B: 175}, 100, 25, false))

	connections.Connect(toggle, led)

	for _, connection := range connections.Data {
		fmt.Println(connection.From().Name(), "->", connection.To().Name())
	}
}
