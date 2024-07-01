package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/apis/data/msgs"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketPlayOutTabInfo struct {
	Header msgs.Message
	Footer msgs.Message
}

func (p *PacketPlayOutTabInfo) UUID() int32 {
	return 71
}

func (p *PacketPlayOutTabInfo) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushTxt(p.Header.AsJson())
	writer.PushTxt(p.Footer.AsJson())
}
