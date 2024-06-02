package build

import (
	"cmp"
	"errors"
	"slices"
	"strconv"

	"github.com/nameless9000/cm2go/block"
)

// converts a block to its compiled string representation
func blockToString(block *block.Base) string {
	blockString := strconv.FormatUint(uint64(block.Id()), 10)

	blockString += ","
	if block.State {
		blockString += "1"
	}
	blockString += ","
	blockString += strconv.FormatFloat(float64(block.Offset.X), 'g', 8, 32)
	blockString += ","
	blockString += strconv.FormatFloat(float64(block.Offset.Y), 'g', 8, 32)
	blockString += ","
	blockString += strconv.FormatFloat(float64(block.Offset.Z), 'g', 8, 32)
	blockString += ","

	count := 0
	for _, property := range block.Properties() {
		blockString += strconv.FormatInt(int64(property), 10) + "+"
		count++
	}

	if count != 0 {
		return blockString[:len(blockString)-1]
	}

	return blockString
}

type CompileFlags struct {
}

// standard collection compilation
func Compile(collectionList []block.Collection, flags CompileFlags) (output string, err error) {
	var blockOutput string
	var connectionOutput string

	var allConnections []*block.Connection = make([]*block.Connection, 0)
	var allBlocks []*block.Base = make([]*block.Base, 0)

	var usageCount map[*block.Base]uint32 = make(map[*block.Base]uint32)

	// populate block and connections
	for _, blocks := range collectionList {
		for _, connection := range blocks.Connections {
			usageCount[connection.From]++
			usageCount[connection.To]++

			allConnections = append(allConnections, connection)
		}

		for _, block := range blocks.Blocks {
			block.Offset.X += blocks.Position.X
			block.Offset.Y += blocks.Position.Y
			block.Offset.Z += blocks.Position.Z

			allBlocks = append(allBlocks, block)
		}
	}

	// sort blocks by most used
	slices.SortFunc(allBlocks, func(a, b *block.Base) int {
		return cmp.Compare(usageCount[a], usageCount[b])
	})

	// convert them to strings
	var blockIndexes map[*block.Base]uint32 = make(map[*block.Base]uint32)

	for count, block := range allBlocks {
		blockIndexes[block] = uint32(count) + 1

		blockOutput += blockToString(block) + ";"
	}

	for _, connection := range allConnections {
		from := strconv.FormatUint(uint64(blockIndexes[connection.From]), 10)
		to := strconv.FormatUint(uint64(blockIndexes[connection.To]), 10)

		connectionOutput += from + "," + to + ";"
	}

	// remove last semicolon and return
	if connectionOutput != "" {
		connectionOutput = connectionOutput[:len(connectionOutput)-1]
	}

	output = blockOutput[:len(blockOutput)-1] + "?" + connectionOutput + "??"
	return
}

// Fast, O(blocks+connections) compilation for debugging and testing
func FastCompile(collectionList []block.Collection) (output string, err error) {
	var blockOutput string
	var connectionOutput string

	var blockIndexes map[*block.Base]uint32 = make(map[*block.Base]uint32)
	var blockCount uint32 = 1

	// compile
	for _, blocks := range collectionList {
		for _, block := range blocks.Blocks {
			block.Offset.X += blocks.Position.X
			block.Offset.Y += blocks.Position.Y
			block.Offset.Z += blocks.Position.Z

			blockIndexes[block] = blockCount
			blockCount++

			blockOutput += blockToString(block) + ";"
		}
	}
	// outside to allow for cross-collection connecting
	for _, blocks := range collectionList {
		for _, connection := range blocks.Connections {
			from := strconv.FormatUint(uint64(blockIndexes[connection.From]), 10)
			to := strconv.FormatUint(uint64(blockIndexes[connection.To]), 10)

			connectionOutput += from + "," + to + ";"
		}
	}

	if blockCount == 1 {
		err = errors.New("at least 1 block is required")
		return
	}

	if connectionOutput != "" {
		connectionOutput = connectionOutput[:len(connectionOutput)-1]
	}

	output = blockOutput[:len(blockOutput)-1] + "?" + connectionOutput + "??"
	return
}
