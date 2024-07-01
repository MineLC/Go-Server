package client

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

type COMMAND int32

const (
	PERFORM_RESPAWN COMMAND = iota
	REQUEST_STATS
	OPEN_INVENTORY_ACHIEVEMENT
)

type PacketPlayInClientCommand struct {
	Action COMMAND
}

func (p *PacketPlayInClientCommand) UUID() int32 {
	return 22
}

func (p *PacketPlayInClientCommand) Pull(reader buff.Buffer, conn base.Connection) {
	p.Action = COMMAND(reader.PullVrI())
}
