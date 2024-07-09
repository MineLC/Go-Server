package impl

import (
	"time"

	"github.com/madflojo/tasks"
	"github.com/minelc/go-server/api"
	"github.com/minelc/go-server/api/cmd"
	"github.com/minelc/go-server/api/data/chat"
	"github.com/minelc/go-server/api/ents"
	"github.com/minelc/go-server/api/network"
	"github.com/minelc/go-server/impl/cmds"

	play "github.com/minelc/go-server/impl/network/server/play"
	srv_tasks "github.com/minelc/go-server/impl/tasks"
)

type server struct {
	players   map[network.Connection]*ents.Player
	scheduler *tasks.Scheduler
	cmd       *cmd.CommandManager
	stop      chan bool
}

func (s *server) GetPlayer(conn network.Connection) *ents.Player {
	return s.players[conn]
}
func (s *server) AddPlayer(conn *network.Connection, player *ents.Player) {
	s.players[*conn] = player
}

func (s *server) Disconnect(conn network.Connection) {
	delete(s.players, conn)
}

func (s *server) GetScheduler() *tasks.Scheduler {
	return s.scheduler
}
func (s *server) GetCommandManager() *cmd.CommandManager {
	return s.cmd
}

func (s *server) Stop() {
	s.scheduler.Stop()
	s.cmd = nil

	for conn := range s.players {
		conn.SendPacket(&play.PacketPlayOutKickDisconnect{Message: *chat.New(chat.Translate("&aStopping server..."))})
		conn.Stop()
		s.stop <- true
	}
}

func Start(stop chan bool) {
	server := server{
		players:   make(map[network.Connection]*ents.Player, 10),
		scheduler: tasks.New(),
		cmd:       cmds.Load(),
		stop:      stop,
	}
	api.SetServer(&server)

	server.scheduler.Add(&tasks.Task{
		Interval: time.Duration(3 * time.Second),
		TaskFunc: func() error { return srv_tasks.KeepAlive(&server.players) },
	})
}
