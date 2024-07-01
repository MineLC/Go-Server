package client

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/apis/uuid"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketPlayInSpectate struct {
	Uuid uuid.UUID
}

func (p *PacketPlayInSpectate) UUID() int32 {
	return 24
}

func (p *PacketPlayInSpectate) Pull(reader buff.Buffer, conn base.Connection) {
	p.Uuid = reader.PullUID()
}
