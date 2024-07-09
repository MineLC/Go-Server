package server

import (
	"github.com/minelc/go-server/api/data/chat"
	"github.com/minelc/go-server/api/network"
)

type PacketPlayOutTabInfo struct {
	Header chat.Message
	Footer chat.Message
}

func (p *PacketPlayOutTabInfo) UUID() int32 {
	return 71
}

func (p *PacketPlayOutTabInfo) Push(writer network.Buffer) {
	writer.PushTxt(p.Header.AsJson())
	writer.PushTxt(p.Footer.AsJson())
}
