package client

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
	"github.com/golangmc/minecraft-server/impl/data/client"
)

type PacketPlayInAbilities struct {
	Abilities   client.PlayerAbilities
	FlightSpeed float32
	GroundSpeed float32
}

func (p *PacketPlayInAbilities) UUID() int32 {
	return 19
}

func (p *PacketPlayInAbilities) Pull(reader buff.Buffer, conn base.Connection) {
	abilities := client.PlayerAbilities{}
	abilities.Pull(reader)

	p.Abilities = abilities

	p.FlightSpeed = reader.PullF32()
	p.GroundSpeed = reader.PullF32()
}
