package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
	"github.com/golangmc/minecraft-server/impl/data/plugin"
)

type PacketPlayOutCustomPayload struct {
	Message plugin.Message
	Buffer  buff.Buffer
}

func (p *PacketPlayOutCustomPayload) UUID() int32 {
	return 63
}

func (p *PacketPlayOutCustomPayload) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushTxt(p.Message.Chan())
	writer.PushUAS(p.Buffer.UAS(), false)
}
