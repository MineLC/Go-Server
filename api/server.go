package api

import (
	"github.com/madflojo/tasks"
	"github.com/minelc/go-server/api/cmd"
	"github.com/minelc/go-server/api/ents"
	"github.com/minelc/go-server/api/network"
)

type Server interface {
	GetPlayer(network.Connection) *ents.Player
	Disconnect(network.Connection)
	GetScheduler() *tasks.Scheduler
	AddPlayer(conn *network.Connection, player *ents.Player)
	GetCommandManager() *cmd.CommandManager
	Broadcast(messages ...string)
	GetConsole() ents.Console
	GetMspt() Mspt
	Stop()
}

var server Server

func GetServer() Server {
	return server
}

func SetServer(newServer Server) {
	if server == nil {
		server = newServer
	}
}
