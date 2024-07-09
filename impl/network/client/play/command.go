package play

import (
	"github.com/minelc/go-server/api/network"
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

func (p *PacketPlayInClientCommand) Pull(reader network.Buffer) {
	p.Action = COMMAND(reader.PullVrI())
}

func (p *PacketPlayInClientCommand) Handle(c *network.Connection) {
	println(p.Action)
}
