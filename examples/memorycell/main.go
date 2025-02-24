package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ko6bxl/cm2go/block"
	"github.com/ko6bxl/cm2go/build"
	"github.com/ko6bxl/cm2go/memory"
)

func register(bits uint32) string {
	collection, register := memory.NewRegister(bits)
	collection.Position.Z = 1

	// inputs
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

	// compile and return
	output, err := build.FastCompile([]block.Collection{collection, inputCollection})
	if err != nil {
		panic(err)
	}

	return output
}

func main() {
	println("Enter amount of bits: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	bits, err := strconv.ParseUint(strings.TrimSpace(input), 10, 32)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(register(uint32(bits)))
}
