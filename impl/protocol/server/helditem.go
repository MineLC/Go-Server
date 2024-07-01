package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
	"github.com/golangmc/minecraft-server/impl/data/client"
)

type PacketPlayOutHeldItemChange struct {
	Slot client.HotBarSlot
}

func (p *PacketPlayOutHeldItemChange) UUID() int32 {
	return 9
}

func (p *PacketPlayOutHeldItemChange) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushByt(byte(p.Slot))
}
