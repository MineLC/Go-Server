package client

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketPlayInKeepAlive struct {
	KeepAliveID int32
}

func (p *PacketPlayInKeepAlive) UUID() int32 {
	return 0
}

func (p *PacketPlayInKeepAlive) Pull(reader buff.Buffer, conn base.Connection) {
	p.KeepAliveID = reader.PullI32()
}
