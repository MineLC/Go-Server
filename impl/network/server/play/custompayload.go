package server

import (
	"github.com/minelc/go-server/api/network"
)

type PacketPlayOutCustomPayload struct {
	Message string
	Buffer  network.Buffer
}

func (p *PacketPlayOutCustomPayload) UUID() int32 {
	return 63
}

func (p *PacketPlayOutCustomPayload) Push(writer network.Buffer) {
	writer.PushTxt(p.Message)
	writer.PushUAS(p.Buffer.UAS(), false)
}
