package client

import (
	"time"

	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketPlayInArmAnimation struct {
	TimeStamp int64
}

func (p *PacketPlayInArmAnimation) UUID() int32 {
	return 10
}

func (p *PacketPlayInArmAnimation) Pull(reader buff.Buffer, conn base.Connection) {
	p.TimeStamp = time.Now().UnixNano() / 1e6
}
