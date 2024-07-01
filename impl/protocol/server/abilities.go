package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
	"github.com/golangmc/minecraft-server/impl/data/client"
)

type PacketPlayOutAbilities struct {
	Abilities   client.PlayerAbilities
	FlyingSpeed float32
	FieldOfView float32
}

func (p *PacketPlayOutAbilities) UUID() int32 {
	return 57
}

func (p *PacketPlayOutAbilities) Push(writer buff.Buffer, conn base.Connection) {
	p.Abilities.Push(writer)

	writer.PushF32(p.FlyingSpeed)
	writer.PushF32(p.FieldOfView)
}
