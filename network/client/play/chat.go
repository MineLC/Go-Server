package play

import (
	"strings"

	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/client/play"
)

func HandleChat(c *network.Connection, packet network.PacketI) {
	p := packet.(*play.PacketPlayInChatMessage)

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

func HandleTab(c *network.Connection, packet network.PacketI) {
	p := packet.(*play.PacketPlayInTabComplete)
	if p.Message[0] != '/' {
		return
	}
	split := strings.Split(p.Message, " ")
	prefix := split[0]
	cmd := api.GetServer().GetCommandManager().Get(prefix[1:])
	if cmd == nil {
		return
	}
	player := api.GetServer().GetPlayer(*c)

	if len(split) > 1 {
		split = split[1:]
	} else {
		split = nil
	}

	cmd.Tab(*player, split)
}
