package play

import (
	"github.com/minelc/go-server/api/network"
)

type PacketPlayInChatMessage struct {
	Message string
}

func (p *PacketPlayInChatMessage) UUID() int32 {
	return 1
}

func (p *PacketPlayInChatMessage) Pull(reader network.Buffer) {
	p.Message = reader.PullTxt()
	if len(p.Message) > 100 {
		p.Message = p.Message[0:99]
	}
}

func (p *PacketPlayInChatMessage) Handle(c *network.Connection) {
	println(p.Message)
}

/*
On tabcomplete a command in chat or command block
*/
type PacketPlayInTabComplete struct {
	Message string
	// Add blockposition of commandblock
}

func (p *PacketPlayInTabComplete) UUID() int32 {
	return 20
}

func (p *PacketPlayInTabComplete) Pull(reader network.Buffer) {
	p.Message = reader.PullTxt()
	/*
		if blockCommand := reader.PullBit(); blockCommand {
			Command block ;v (elpapu dice que no)
		}
	*/
}

func (p *PacketPlayInTabComplete) Handle(c *network.Connection) {
	println(p.Message)
}
