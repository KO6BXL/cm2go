package build

import (
	"cmp"
	"errors"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/ko6bxl/cm2go/block"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		println(name, "took", time.Since(start).Microseconds(), "microseconds\n")
	}
}

// converts a block to its compiled string representation
func writeBlockString(block *block.Base, stringBuilder *strings.Builder) {
	stringBuilder.WriteString(strconv.FormatUint(uint64(block.Id()), 10))

	stringBuilder.WriteRune(',')
	if block.State {
		stringBuilder.WriteRune('1')
	}
	stringBuilder.WriteRune(',')
	stringBuilder.WriteString(strconv.FormatFloat(float64(block.Offset.X), 'g', 8, 32))
	stringBuilder.WriteRune(',')
	stringBuilder.WriteString(strconv.FormatFloat(float64(block.Offset.Y), 'g', 8, 32))
	stringBuilder.WriteRune(',')
	stringBuilder.WriteString(strconv.FormatFloat(float64(block.Offset.Z), 'g', 8, 32))
	stringBuilder.WriteRune(',')

	properties := len(block.Properties())
	for count, property := range block.Properties() {
		stringBuilder.WriteString(strconv.FormatInt(int64(property), 10))

		if count != properties-1 {
			stringBuilder.WriteRune('+')
		}
	}
}

// standard collection compilation
func Compile(collectionList []block.Collection) (output string, err error) {
	defer timer("Compile")()
	var stringBuilder strings.Builder

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

	blockLen := len(allBlocks)
	for count, block := range allBlocks {
		blockIndexes[block] = uint32(count) + 1

		writeBlockString(block, &stringBuilder)
		if count != blockLen-1 {
			stringBuilder.WriteRune(';')
		}
	}
	println("Block count:", blockLen)

	stringBuilder.WriteString("?")

	connectionLen := len(allConnections)
	for count, connection := range allConnections {
		from := strconv.FormatUint(uint64(blockIndexes[connection.From]), 10)
		to := strconv.FormatUint(uint64(blockIndexes[connection.To]), 10)

		stringBuilder.WriteString(from)
		stringBuilder.WriteRune(',')
		stringBuilder.WriteString(to)
		if count != connectionLen-1 {
			stringBuilder.WriteRune(';')
		}
	}
	println("Connection count:", connectionLen)

	stringBuilder.WriteString("??")

	output = stringBuilder.String()
	return
}

// Fast, O(blocks+connections) compilation for debugging and testing
func FastCompile(collectionList []block.Collection) (output string, err error) {
	defer timer("Compile")()
	var stringBuilder strings.Builder

	var blockIndexes map[*block.Base]uint32 = make(map[*block.Base]uint32)
	var blockCount uint32 = 1

	// compile
	for collectionCount, blocks := range collectionList {
		isLast := collectionCount == len(collectionList)-1
		blockLen := len(blocks.Blocks)
		for count, block := range blocks.Blocks {
			block.Offset.X += blocks.Position.X
			block.Offset.Y += blocks.Position.Y
			block.Offset.Z += blocks.Position.Z

			blockIndexes[block] = blockCount
			blockCount++

			writeBlockString(block, &stringBuilder)
			if !isLast || count != blockLen-1 {
				stringBuilder.WriteRune(';')
			}
		}
	}
	println("Block count:", blockCount-1)

	stringBuilder.WriteString("?")

	// outside to allow for cross-collection connecting
	connectionCount := 0
	for collectionCount, blocks := range collectionList {
		isLast := collectionCount == len(collectionList)-1
		connectionLen := len(blocks.Connections)
		for count, connection := range blocks.Connections {
			from := strconv.FormatUint(uint64(blockIndexes[connection.From]), 10)
			to := strconv.FormatUint(uint64(blockIndexes[connection.To]), 10)

			stringBuilder.WriteString(from)
			stringBuilder.WriteRune(',')
			stringBuilder.WriteString(to)
			if !isLast || count != connectionLen-1 {
				stringBuilder.WriteRune(';')
			}
			connectionCount++
		}
	}
	println("Connection count:", connectionCount)

	if blockCount == 1 {
		err = errors.New("at least 1 block is required")
		return
	}

	stringBuilder.WriteString("??")

	output = stringBuilder.String()
	return
}

func MemCompile(data []uint8) (string, error) {

	hextable := []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
	}
	var out string
	if len(data) > 4096 {
		return out, errors.New("Too much data for massMemory")
	}

	var popbits int
	for i, x := range data {
		var hex string
		var lb uint8
		var hb uint8

		hb = x >> 4
		lb = (x << 4) >> 4
		hex = hex + hextable[hb]
		hex = hex + hextable[lb]
		out = out + hex
		popbits = i
	}

	for range 4096 - popbits {
		out = out + "00"
	}

	return out, nil
}
