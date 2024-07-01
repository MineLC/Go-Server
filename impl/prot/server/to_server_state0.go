package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketHandshakingInSetProtocol struct {
	version int32

	host string
	port uint16

	State base.PacketState
}

func (p *PacketHandshakingInSetProtocol) UUID() int32 {
	return 0
}

func (p *PacketHandshakingInSetProtocol) Pull(reader buff.Buffer, conn base.Connection) {
	p.version = reader.PullVrI()

	p.host = reader.PullTxt()
	p.port = reader.PullU16()

	p.State = base.PacketStateValueOf(int(reader.PullVrI()))
}
