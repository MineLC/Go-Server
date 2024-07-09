package server

import (
	"github.com/minelc/go-server/api/data/player"
	"github.com/minelc/go-server/api/ents"
	"github.com/minelc/go-server/api/network"
)

type PacketPlayOutEntityMetadata struct {
	Entity ents.Entity
}

func (p *PacketPlayOutEntityMetadata) UUID() int32 {
	return 28
}

func (p *PacketPlayOutEntityMetadata) Push(writer network.Buffer) {
	writer.PushVrI(int32(p.Entity.EntityUUID())) // questionable...

	// only supporting player metadata for now
	_, ok := p.Entity.(ents.Player)
	if ok {

		writer.PushByt(16) // index | displayed skin parts
		writer.PushVrI(0)  // type | byte

		skin := player.SkinParts{
			Cape: true,
			Head: true,
			Body: true,
			ArmL: true,
			ArmR: true,
			LegL: true,
			LegR: true,
		}

		skin.Push(writer)
	}
}
