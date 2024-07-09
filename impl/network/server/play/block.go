package server

import (
	"github.com/minelc/go-server/api/network"
	"github.com/minelc/go-server/api/data"
)

type PacketPlayOutBlockChange struct {
	Pos     data.PositionI
	BlockID data.Block
	Data    int32
}

func (p *PacketPlayOutBlockChange) UUID() int32 {
	return 35
}

func (p *PacketPlayOutBlockChange) Push(writer network.Buffer) {
	writer.PushPos(p.Pos)
	writer.PushVrI(int32(p.BlockID) + p.Data<<12)
}
