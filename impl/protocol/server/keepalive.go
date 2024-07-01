package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketPlayOutKeepAlive struct {
	KeepAliveID int32
}

func (p *PacketPlayOutKeepAlive) UUID() int32 {
	return 0
}

func (p *PacketPlayOutKeepAlive) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushVrI(p.KeepAliveID)
}
