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

func (connections *Collection) Connect(from *Base, to *Base) {
	connection := new(Connection)
	connection.From = from
	connection.To = to

	connections.Connections = append(connections.Connections, connection)
}
