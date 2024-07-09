package cmds

import (
	"github.com/minelc/go-server/api"
	"github.com/minelc/go-server/api/cmd"
	"github.com/minelc/go-server/api/ents"
)

func Load() *cmd.CommandManager {
	manager := cmd.NewCommandManager()

	manager.AddStruct(cmd.StructCommand{
		Execute: func(sender ents.Sender, args []string) {
			sender.SendMsg(
				" ",
				" &b&lGo Server &f- &71.8",
				" ",
				" &fFollow the project on github:",
				" &bhttps://github.com/MineLC/Go-Server",
			)
		},
	}, "version")

	manager.AddStruct(cmd.StructCommand{
		Execute: func(sender ents.Sender, args []string) {
			api.GetServer().Stop()
		},
	}, "stop")

	return &manager
}
