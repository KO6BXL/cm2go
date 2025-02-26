package block

type Connection struct {
	To   *Base
	From *Base
}

type Collection struct {
	Blocks      []*Base
	Connections []*Connection
	Position    Offset
}

// Appends a block to the collection and returns the block
func (collection *Collection) Append(block *Base) *Base {
	collection.Blocks = append(collection.Blocks, block)
	return block
}

func (collection *Collection) Connect(from *Base, to ...*Base) {
	for _, connectTo := range to {
		connection := new(Connection)
		connection.From = from
		connection.To = connectTo

		collection.Connections = append(collection.Connections, connection)
	}
}

// Merges the collection's blocks and connections
func (collection Collection) Merge(with ...Collection) {
	for _, toMerge := range with {
		for _, block := range toMerge.Blocks {
			block.Offset.X += toMerge.Position.X
			block.Offset.Y += toMerge.Position.Y
			block.Offset.Z += toMerge.Position.Z
		}

		collection.Blocks = append(collection.Blocks, toMerge.Blocks...)
		collection.Connections = append(collection.Connections, toMerge.Connections...)
	}
}
