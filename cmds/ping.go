package cmds

import (
	"fmt"

	"github.com/minelc/go-server-api/ents"
)

func PingCMD(sender ents.Sender, args []string) {
	player, ok := sender.(ents.Player)

	if !ok {
		sender.SendMsgColor("&cOnly players")
		return
	}
	player.SendMsgColor(fmt.Sprint("&aPing: ", player.GetPing(), "ms"))
}
