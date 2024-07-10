package cmds

import (
	"fmt"

	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/cmd"
	"github.com/minelc/go-server-api/ents"
)

func Load() *cmd.CommandManager {
	manager := cmd.NewCommandManager()

	manager.AddStruct(cmd.StructCommand{
		Execute: func(sender ents.Sender, args []string) {
			sender.SendMsgColor(
				" ",
				" &b&lGo Server &f- &71.8",
				" ",
				" &fFollow the project on github:",
				" &bhttps://github.com/MineLC/Go-Server",
			)
		},
	}, "version")

	manager.AddStruct(cmd.StructCommand{
		Execute: PingCMD,
	}, "ping")

	manager.AddStruct(cmd.StructCommand{
		Execute: func(sender ents.Sender, args []string) {
			mspt := api.GetServer().GetMspt()
			sender.SendMsgColor(
				"  &b&l MSPT: &7(milliseconds per tick): ",
				"  &7Last 20s: ",
				fmt.Sprint("       &fMax: &f", mspt.GetMax()),
				fmt.Sprint("       &fMed: &b", mspt.GetPromedium()),
				fmt.Sprint("       &fMin: &3", mspt.GetMin()),
			)
		},
	}, "mspt")

	manager.AddStruct(cmd.StructCommand{
		Execute: func(sender ents.Sender, args []string) {
			api.GetServer().Stop()
		},
	}, "stop")

	return &manager
}
