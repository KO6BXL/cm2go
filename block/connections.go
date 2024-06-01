package block

type Connection struct {
	to   *Base
	from *Base
}

func (connection Connection) To() *Base {
	return connection.to
}

func (connection Connection) From() *Base {
	return connection.from
}

type Connections struct {
	Data []*Connection
}

func (connections *Connections) Connect(from *Base, to *Base) {
	connection := new(Connection)
	connection.from = from
	connection.to = to

	connections.Data = append(connections.Data, connection)
}
