package play

import (
	"time"

	"github.com/minelc/go-server/api/network"
)

type PacketPlayInArmAnimation struct {
	TimeStamp int64
}

func (p *PacketPlayInArmAnimation) UUID() int32 {
	return 10
}

func (p *PacketPlayInArmAnimation) Pull(reader network.Buffer) {
	p.TimeStamp = time.Now().UnixNano() / 1e6
}

func (p *PacketPlayInArmAnimation) Handle(c *network.Connection) {

}
