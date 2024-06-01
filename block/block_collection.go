package block

type Collection struct {
	Blocks   []*Base
	Position Offset
}

// Appends a block to the collection and returns the block
func (collection *Collection) Append(block *Base) *Base {
	collection.Blocks = append(collection.Blocks, block)
	return block
}
