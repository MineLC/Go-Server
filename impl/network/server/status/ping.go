package status

import "github.com/minelc/go-server/api/network"

type PacketStatusOPing struct {
	Ping int64
}

func (p *PacketStatusOPing) Push(out network.Buffer) {
	out.PushI64(p.Ping)
}

func (p *PacketStatusOPing) UUID() int32 {
	return 1
}
