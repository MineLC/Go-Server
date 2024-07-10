package play

import (
	"strings"

	"github.com/minelc/go-server/api"
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
	uplayer := api.GetServer().GetPlayer(*c)
	if uplayer == nil {
		return
	}
	player := *uplayer
	if p.Message[0] != '/' {
		api.GetServer().Broadcast(player.GetProfile().Name + " : " + p.Message)
		return
	}

	split := strings.Split(p.Message, " ")
	length := len(split)
	if length < 1 {
		return
	}

	prefix := split[0]
	cmd := api.GetServer().GetCommandManager().Get(prefix)
	if cmd == nil {
		player.SendMsgColor("&cThis command don't exist")
		return
	}

	if length > 1 {
		split = split[1:]
	} else {
		split = nil
	}
	cmd.Execute(player, split)
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
