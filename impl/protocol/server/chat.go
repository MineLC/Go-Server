package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/apis/data/msgs"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketOutChatMessage struct {
	Message         msgs.Message
	MessagePosition msgs.MessagePosition
}

func (p *PacketOutChatMessage) UUID() int32 {
	return 2
}

func (p *PacketOutChatMessage) Push(writer buff.Buffer, conn base.Connection) {
	message := p.Message

	if p.MessagePosition == msgs.HotBarText {
		message = *msgs.New(message.AsText())
	}
	writer.PushTxt(message.AsJson())
	writer.PushByt(byte(p.MessagePosition))
}
