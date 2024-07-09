package status

import (
	"github.com/minelc/go-server/api/network"
)

// Send to request the motd. But we use handshake, so this packet is useless
type PacketIRequest struct{}

func (p *PacketIRequest) Pull(reader network.Buffer) {}

func (p *PacketIRequest) UUID() int32 {
	return 0
}
func (p *PacketIRequest) Handle(conn *network.Connection) {}
