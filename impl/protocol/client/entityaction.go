package client

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PLAYER_ACTION int32

const (
	START_SNEAKING PLAYER_ACTION = iota
	STOP_SNEAKING
	STOP_SLEEPING
	START_SPRINTING
	STOP_SPRINTING
	RIDING_JUMP
	OPEN_INVENTORY
)

type PacketPlayInEntityAction struct {
	EntityID int32
	Action   PLAYER_ACTION
	C        int32
}

func (p *PacketPlayInEntityAction) UUID() int32 {
	return 11
}

func (p *PacketPlayInEntityAction) Pull(reader buff.Buffer, conn base.Connection) {
	p.EntityID = reader.PullVrI()
	p.Action = PLAYER_ACTION(reader.PullVrI())
	p.C = reader.PullVrI()
}
