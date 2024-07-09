package play

import (
	"github.com/minelc/go-server/api/data"
	"github.com/minelc/go-server/api/network"
)

type PacketPlayInSpectate struct {
	Uuid data.UUID
}

func (p *PacketPlayInSpectate) UUID() int32 {
	return 24
}

func (p *PacketPlayInSpectate) Pull(reader network.Buffer) {
	p.Uuid = reader.PullUID()
}
func (p *PacketPlayInSpectate) Handle(c *network.Connection) {

}
