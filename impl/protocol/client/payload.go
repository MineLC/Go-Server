package client

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketPlayInCustomPayload struct {
	Prefix string
	Buffer []int8
}

func (p *PacketPlayInCustomPayload) UUID() int32 {
	return 23
}

func (p *PacketPlayInCustomPayload) Pull(reader buff.Buffer, conn base.Connection) {
	p.Prefix = reader.PullTxt()
	p.Buffer = reader.PullSAS()
}
