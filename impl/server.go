package impl

import (
	"time"

	"github.com/madflojo/tasks"
	"github.com/minelc/go-server/api"
	"github.com/minelc/go-server/api/ents"
	"github.com/minelc/go-server/api/network"
	srv_tasks "github.com/minelc/go-server/impl/tasks"
)

type server struct {
	players   map[network.Connection]*ents.Player
	scheduler *tasks.Scheduler
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

func Start() {
	server := server{
		players:   make(map[network.Connection]*ents.Player, 10),
		scheduler: tasks.New(),
	}
	api.SetServer(&server)

	server.scheduler.Add(&tasks.Task{
		Interval: time.Duration(3 * time.Second),
		TaskFunc: func() error { return srv_tasks.KeepAlive(&server.players) },
	})
}
