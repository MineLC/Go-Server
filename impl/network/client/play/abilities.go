package play

import (
	"github.com/minelc/go-server/api/data/player"
	"github.com/minelc/go-server/api/network"
)

type PacketPlayInAbilities struct {
	Abilities   player.PlayerAbilities
	FlightSpeed float32
	GroundSpeed float32
}

func (p *PacketPlayInAbilities) UUID() int32 {
	return 19
}

func (p *PacketPlayInAbilities) Pull(reader network.Buffer) {
	abilities := player.PlayerAbilities{}
	abilities.Pull(reader)

	p.Abilities = abilities

	p.FlightSpeed = reader.PullF32()
	p.GroundSpeed = reader.PullF32()
}

func (p *PacketPlayInAbilities) Handle(c *network.Connection) {

}
