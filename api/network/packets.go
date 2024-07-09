package network

type PacketI interface {
	Pull(in Buffer)
	UUID() int32
	Handle(conn *Connection)
}

type PacketO interface {
	UUID() int32
	Push(out Buffer)
}
