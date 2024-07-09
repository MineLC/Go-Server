package server

import (
	"github.com/minelc/go-server/api/data/player"
	"github.com/minelc/go-server/api/network"
)

type PacketPlayOutAbilities struct {
	Abilities   player.PlayerAbilities
	FlyingSpeed float32
	FieldOfView float32
}

func (p *PacketPlayOutAbilities) UUID() int32 {
	return 57
}

func (p *PacketPlayOutAbilities) Push(writer network.Buffer) {
	p.Abilities.Push(writer)

	writer.PushF32(p.FlyingSpeed)
	writer.PushF32(p.FieldOfView)
}
