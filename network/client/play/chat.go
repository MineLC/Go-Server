package play

import (
	"strings"

	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/client/play"
	"github.com/minelc/go-server-api/plugin"
	"github.com/minelc/go-server-api/plugin/events"
)

func HandleChat(c network.Connection, packet network.PacketI) {
	p := packet.(*play.PacketPlayInChatMessage)

	player := api.GetServer().GetPlayer(c)
	if player == nil {
		return
	}
	if p.Message[0] != '/' {
		event := events.PlayerChat{Cancel: false, Player: player, Message: p.Message}
		api.GetServer().GetPluginManager().CallEvent(event, plugin.Chat)
		if !event.Cancel {
			api.GetServer().Broadcast(player.GetProfile().Name + " : " + p.Message)
		}
		return
	}

	split := strings.Split(p.Message, " ")
	length := len(split)
	if length < 1 {
		return
	}

	prefix := split[0]
	cmd := api.GetServer().GetPluginManager().GetCommandManager().Get(prefix)
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

func HandleTab(c network.Connection, packet network.PacketI) {
	p := packet.(*play.PacketPlayInTabComplete)
	if p.Message[0] != '/' {
		return
	}
	split := strings.Split(p.Message, " ")
	prefix := split[0]
	cmd := api.GetServer().GetPluginManager().GetCommandManager().Get(prefix[1:])
	if cmd == nil {
		return
	}
	player := api.GetServer().GetPlayer(c)

	if len(split) > 1 {
		split = split[1:]
	} else {
		split = nil
	}

	cmd.Tab(player, split)
}
