package status

import (
	"github.com/minelc/go-server/api/network"
	"github.com/minelc/go-server/impl/network/server/status"
)

type PacketIPing struct {
	Ping int64
}

func (p *PacketIPing) Pull(reader network.Buffer) {
}

func (p *PacketIPing) UUID() int32 {
	return 1
}

func (p *PacketIPing) Handle(conn *network.Connection) {
	(*conn).SendPacket(&status.PacketStatusOPing{Ping: p.Ping})
}
