package status

import "github.com/minelc/go-server/api/network"

type PacketOResponse struct {
	Motd string
}

func (p *PacketOResponse) UUID() int32 {
	return 0
}

func (p *PacketOResponse) Push(writer network.Buffer) {
	writer.PushTxt(p.Motd)
}
