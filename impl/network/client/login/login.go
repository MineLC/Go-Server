package login

import (
	"github.com/minelc/go-server/api/network"
)

type PacketILoginPluginResponse struct {
	Message int32
	Success bool
	OptData []byte
}

func (p *PacketILoginPluginResponse) UUID() int32 {
	return 0x02
}

func (p *PacketILoginPluginResponse) Pull(reader network.Buffer) {
	p.Message = reader.PullVrI()
	p.Success = reader.PullBit()
	p.OptData = reader.UAS()[reader.InI():reader.Len()]
}

func (p *PacketILoginPluginResponse) Handle(conn *network.Connection) {}
