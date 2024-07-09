package server

import (
	"github.com/minelc/go-server/api/data/chat"
	"github.com/minelc/go-server/api/network"
)

type PacketOutChatMessage struct {
	Message         chat.Message
	MessagePosition chat.MessagePosition
}

func (p *PacketOutChatMessage) UUID() int32 {
	return 2
}

func (p *PacketOutChatMessage) Push(writer network.Buffer) {
	message := p.Message

	if p.MessagePosition == chat.HotBarText {
		message = *chat.New(message.AsText())
	}
	writer.PushTxt(message.AsJson())
	writer.PushByt(byte(p.MessagePosition))
}
