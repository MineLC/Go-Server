package play

import (
	"github.com/minelc/go-server/api/network"
)

type ENTITY_ACTION int32

const (
	INTERACT ENTITY_ACTION = iota
	ATTACK
	INTERACT_AT
)

type PacketPlayInUseEntity struct {
	EntityID int32
	Action   ENTITY_ACTION

	VectorX float32
	VectorY float32
	VectorZ float32
}

func (p *PacketPlayInUseEntity) UUID() int32 {
	return 2
}

func (p *PacketPlayInUseEntity) Pull(reader network.Buffer) {
	p.EntityID = reader.PullVrI()
	action := reader.PullVrI()
	p.Action = ENTITY_ACTION(action)
	if p.Action == INTERACT_AT {
		p.VectorX = reader.PullF32()
		p.VectorY = reader.PullF32()
		p.VectorZ = reader.PullF32()
	}
}

func (p *PacketPlayInUseEntity) Handle(c *network.Connection) {

}
