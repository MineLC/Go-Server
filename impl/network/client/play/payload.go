package play

import (
	"github.com/minelc/go-server/api/network"
)

type PacketPlayInCustomPayload struct {
	Prefix string
	Buffer []int8
}

func (p *PacketPlayInCustomPayload) UUID() int32 {
	return 23
}

func (p *PacketPlayInCustomPayload) Pull(reader network.Buffer) {
	p.Prefix = reader.PullTxt()
	p.Buffer = reader.PullSAS()
}

func (p *PacketPlayInCustomPayload) Handle(c *network.Connection) {

}
