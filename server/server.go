package server

import (
	"fmt"
	"time"

	"github.com/madflojo/tasks"
	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/data/chat"
	"github.com/minelc/go-server-api/ents"
	"github.com/minelc/go-server-api/game/world"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/server/play"

	api_plugin "github.com/minelc/go-server-api/plugin"
	"github.com/minelc/go-server/cmds"
	"github.com/minelc/go-server/conf"

	impl_world "github.com/minelc/go-server/game/world"
	impl_net "github.com/minelc/go-server/network"
	plugin "github.com/minelc/go-server/plugin"
	srv_tasks "github.com/minelc/go-server/tasks"
)

type Server struct {
	players map[network.Connection]ents.Player

	console       *Console
	scheduler     *tasks.Scheduler
	pluginManager api_plugin.PluginManager
	mspt          Mspt
	running       bool
	packets       network.PacketManager
	worlds        world.WorldManager
}

func (s *Server) GetWorldManager() world.WorldManager {
	return s.worlds
}

func (s *Server) GetPlayer(conn network.Connection) ents.Player {
	return s.players[conn]
}
func (s *Server) AddPlayer(conn network.Connection, player ents.Player) {
	s.players[conn] = player
}
func (s *Server) Disconnect(conn network.Connection) {
	delete(s.players, conn)
}
func (s *Server) GetPlayers() *map[network.Connection]ents.Player {
	return &s.players
}

func (s *Server) GetScheduler() *tasks.Scheduler {
	return s.scheduler
}
func (s *Server) GetPluginManager() api_plugin.PluginManager {
	return s.pluginManager
}
func (s *Server) GetConsole() ents.Console {
	return s.console
}
func (s *Server) GetMspt() api.Mspt {
	return s.mspt
}
func (s *Server) GetPacketManager() network.PacketManager {
	return s.packets
}

func (s *Server) Broadcast(messages ...string) {
	for _, msg := range messages {
		msgPacket := &play.PacketPlayOutChatMessage{Message: *chat.New(msg)}
		for conn := range s.players {
			conn.SendPacket(msgPacket)
		}
		s.console.SendMsg(msg)
	}
}

func (s *Server) Stop() {
	complete := make(chan bool)
	go plugin.StopPlugins(s, complete)
	<-complete

	s.scheduler.Stop()
	s.pluginManager = nil
	s.console.SendMsgColor("&aStopping server...")

	kickPacket := &play.PacketPlayOutKickDisconnect{Message: *chat.New(chat.Translate("&aStopping server..."))}
	for conn := range s.players {
		conn.SendPacket(kickPacket)
		conn.Stop()
	}
	s.running = false
}

func (s *Server) LoadPlugins() {
	pluginsTime := time.Now().UnixMicro()
	amount := plugin.LoadPlugins(s)
	if amount == 0 {
		s.console.SendMsgColor("&3No plugins found!")
		return
	}
	pluginsTime = time.Now().UnixMicro() - pluginsTime
	s.console.SendMsgColor("&aPlugins loaded: &e" + fmt.Sprint(amount) + " &ain: &e" + fmt.Sprint(pluginsTime) + "&a µs")
}

func Start(conf conf.ServerConfig) *Server {
	if api.GetServer() != nil {
		api.GetServer().Stop()
		return nil
	}

	c := Console{}
	pluginManager := plugin.NewPluginManager(cmds.Load())

	server := Server{
		players:   make(map[network.Connection]ents.Player, 3),
		scheduler: tasks.New(),
		console:   &c,
		mspt: Mspt{
			nextTwentySeconds: time.Now().UnixMilli() + 20_000,
		},
		pluginManager: pluginManager,
		packets:       impl_net.NewDefaultHandler(conf.Network.Compression),
		running:       true,
	}

	api.SetServer(&server)
	go c.start()

	server.scheduler.Add(&tasks.Task{
		Interval: time.Duration(15 * time.Second),
		TaskFunc: func() error { return srv_tasks.KeepAlive(&server.players) },
	})

	worldManager := impl_world.NewWorldManager(conf.Game.DefaultWorld)
	server.worlds = worldManager

	return &server
}

func StartMainLoop(s *Server) {
	var sleepTime time.Duration = 50
	var delayedTicks int32 = 0

	for {
		if !s.running {
			s.console.SendMsgColor("&cServer stopped. Thanks for use Go-Server :)")
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

func executeMain(s *Server) {
	s.console.executePendient()

}
