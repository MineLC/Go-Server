package play

import (
	"time"

	"github.com/minelc/go-server/api"
	"github.com/minelc/go-server/api/network"
	"github.com/minelc/go-server/impl/tasks"
)

type PacketPlayInKeepAlive struct {
	KeepAliveID int32
}

func (p *PacketPlayInKeepAlive) UUID() int32 {
	return 0
}

func (p *PacketPlayInKeepAlive) Pull(reader network.Buffer) {
	p.KeepAliveID = reader.PullI32()
}

func (p *PacketPlayInKeepAlive) Handle(c *network.Connection) {
	player := api.GetServer().GetPlayer(*c)
	if player != nil {
		now := time.Now().Unix()
		(*player).SetKeepAlive(now)
		(*player).SetPing(now, tasks.GetLastAlive())
	}
}
