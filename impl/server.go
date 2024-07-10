package impl

import (
	"fmt"
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
	console   *Console
	scheduler *tasks.Scheduler
	cmd       *cmd.CommandManager
	mspt      Mspt
	running   bool
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
func (s *server) GetConsole() ents.Console {
	return s.console
}
func (s *server) GetMspt() api.Mspt {
	return s.mspt
}

func (s *server) Broadcast(messages ...string) {
	for _, msg := range messages {
		msgPacket := &play.PacketPlayOutChatMessage{Message: *chat.New(msg)}
		for conn := range s.players {
			conn.SendPacket(msgPacket)
		}
		s.console.SendMsg(msg)
	}
}

func (s *server) Stop() {
	s.scheduler.Stop()
	s.cmd = nil
	s.console.SendMsgColor("&aStopping server...")

	kickPacket := &play.PacketPlayOutKickDisconnect{Message: *chat.New(chat.Translate("&aStopping server..."))}
	for conn := range s.players {
		conn.SendPacket(kickPacket)
		conn.Stop()
	}
	s.running = false
}

func Start() {
	if api.GetServer() != nil {
		api.GetServer().Stop()
	}

	c := Console{}
	server := server{
		players:   make(map[network.Connection]*ents.Player, 10),
		scheduler: tasks.New(),
		cmd:       cmds.Load(),
		console:   &c,
		mspt: Mspt{
			max:               0,
			promedium:         0,
			min:               0,
			elapseTicks:       0,
			nextTwentySeconds: time.Now().UnixMilli() + 20_000,
		},
		running: true,
	}

	api.SetServer(&server)
	go c.start()

	server.scheduler.Add(&tasks.Task{
		Interval: time.Duration(15 * time.Second),
		TaskFunc: func() error { return srv_tasks.KeepAlive(&server.players) },
	})

	startMainLoop(&server)
}

func startMainLoop(s *server) {
	var sleepTime time.Duration = 50
	var delayedTicks int32 = 0

	for {
		if !s.running {
			return
		}
		if delayedTicks != 0 {
			delayedTicks--
		} else {
			time.Sleep(sleepTime * time.Millisecond)
		}

		s.mspt.elapseTicks++
		startTime := time.Now().UnixMilli()

		executeMain(s) // Execute tasks, console, etc

		mspt := s.mspt.Handle(startTime)

		// How works?: If a tick take x milliseconds, the main thread sleep less time
		// To mantain the sleep of 50 milliseconds (a complete tick)
		millisDelayed := 50 - mspt

		// If the tick take more than 50ms is a delayed tick
		// Example: a tick take 100ms, there are 1 ticks delayed
		if millisDelayed < 0 {
			delayedTicks = int32((millisDelayed - 1) / 50)
			s.console.SendMsgColor(fmt.Sprint("A tick take more than 50ms. Delayed ticks: ", delayedTicks))
			continue
		}
		sleepTime = time.Duration(millisDelayed)
	}
}

func executeMain(s *server) {
	s.console.executePendient()
}
