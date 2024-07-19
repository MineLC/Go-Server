package join

import (
	"github.com/minelc/go-server-api/ents"
	"github.com/minelc/go-server-api/network"
)

/*
This packets will be add in the api.
Now, this are for testing and debug
*/
type PacketPlayOutEntityMetadata struct {
	Entity ents.Entity
}

func (p *PacketPlayOutEntityMetadata) UUID() int32 {
	return 0x1C
}

func (p *PacketPlayOutEntityMetadata) Push(writer network.Buffer) {
	writer.PushVrI(int32(p.Entity.EntityUUID()))
	p.Entity.PushMetadata(writer)
	writer.PushByt(127)
}

type PacketPlayOutExperience struct {
	Xp    int32
	Level int32
}

func (p *PacketPlayOutExperience) UUID() int32 {
	return 0x1F
}

func (p *PacketPlayOutExperience) Push(writer network.Buffer) {
	writer.PushF32(float32(1 / p.Xp))
	writer.PushVrI(p.Level)
	writer.PushVrI(p.Xp)
}
