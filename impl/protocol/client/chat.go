package client

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketPlayInChatMessage struct {
	Message string
}

func (p *PacketPlayInChatMessage) UUID() int32 {
	return 1
}

func (p *PacketPlayInChatMessage) Pull(reader buff.Buffer, conn base.Connection) {
	p.Message = reader.PullTxt()
	if len(p.Message) > 100 {
		p.Message = p.Message[0:99]
	}
}

/*
On tabcomplete a command in chat or command block
*/
type PacketPlayInTabComplete struct {
	Message string
	// TODO: Add blockposition of commandblock
}

func (p *PacketPlayInTabComplete) UUID() int32 {
	return 20
}

func (p *PacketPlayInTabComplete) Pull(reader buff.Buffer, conn base.Connection) {
	p.Message = reader.PullTxt()
	if blockCommand := reader.PullBit(); blockCommand {
		// TODO: Command block :(
		print("CommandBlock :(")
	}
}
