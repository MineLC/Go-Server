package server

import (
	"github.com/minelc/go-server/api/data/chat"
	"github.com/minelc/go-server/api/network"
)

type PacketPlayOutKickDisconnect struct {
	Message chat.Message
}

func (p *PacketPlayOutKickDisconnect) UUID() int32 {
	return 64
}

func (p *PacketPlayOutKickDisconnect) Push(writer network.Buffer) {
	writer.PushTxt(p.Message.AsJson())
}
