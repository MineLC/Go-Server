package login

import (
	"github.com/minelc/go-server/api/data/chat"
	"github.com/minelc/go-server/api/network"
)

type PacketODisconnect struct {
	Reason chat.Message
}

func (p *PacketODisconnect) UUID() int32 {
	return 0x00
}

func (p *PacketODisconnect) Push(writer network.Buffer) {
	message := p.Reason

	writer.PushTxt(message.AsJson())
}
