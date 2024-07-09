package server

import (
	"github.com/minelc/go-server/api/data/chat"
	"github.com/minelc/go-server/api/network"
)

type PacketPlayOutChatMessage struct {
	Message         chat.Message
	MessagePosition chat.MessagePosition
}

func (p *PacketPlayOutChatMessage) UUID() int32 {
	return 2
}

func (p *PacketPlayOutChatMessage) Push(writer network.Buffer) {
	message := p.Message

	if p.MessagePosition == chat.HotBarText {
		message = *chat.New(message.AsText())
	}
	writer.PushTxt(message.AsJson())
	writer.PushByt(byte(p.MessagePosition))
}
