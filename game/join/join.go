package join

import (
	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/data/chat"
	"github.com/minelc/go-server-api/data/player"
	"github.com/minelc/go-server-api/ents"
	"github.com/minelc/go-server-api/game"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/server/login"
	"github.com/minelc/go-server-api/network/server/play"
	"github.com/minelc/go-server-api/plugin"
	"github.com/minelc/go-server-api/plugin/events"
	"github.com/minelc/go-server-api/world"
	"github.com/minelc/go-server/conf"
)

var GameOptions conf.Game

func Join(p ents.Player, conn network.Connection) {
	conn.SendPacket(&login.PacketOLoginSuccess{
		PlayerName: p.GetProfile().Name,
		PlayerUUID: p.UUID().String(),
	})
	conn.SetState(network.PLAY)
	api.GetServer().GetPluginManager().CallEvent(events.PlayerJoinEvent{Player: p}, plugin.Join)

	sendPlayPackets(p, conn)

	api.GetServer().AddPlayer(conn, p)
}

func sendPlayPackets(p ents.Player, conn network.Connection) {
	conn.SendPacket(&play.PacketPlayOutLogin{
		EntityID:    int32(p.EntityUUID()),
		Hardcore:    false,
		GameMode:    player.SURVIVAL,
		Dimension:   world.OVERWORLD,
		Difficulty:  game.EASY,
		MaxPlayers:  10,
		LevelType:   world.CUSTOMIZED,
		ReduceDebug: false,
	})

	conn.SendPacket(&play.PacketPlayOutServerDifficulty{
		Difficulty: game.NORMAL,
	})

	world := api.GetServer().GetWorldManager().GetDefaultWorld()
	if world != nil {
		conn.SendPacket(&play.PacketPlayOutBulkChunkData{Chunks: world.GetAllChunks()})
	}

	sendTab(p, conn)
	debugPackets(p, conn)

	conn.SendPacket(&play.PacketPlayOutAbilities{
		Abilities: player.PlayerAbilities{
			Invulnerable: true,
			Flying:       true,
			AllowFlight:  true,
			InstantBuild: false,
		},
		FlyingSpeed: 0.05, // default value
		FieldOfView: 0.1,  // default value
	})
}

func sendTab(p ents.Player, conn network.Connection) {
	addPlayer := play.PacketPlayOutPlayerInfo{Action: play.ADD_PLAYER, Players: []play.PlayerInfoData{
		{
			Profile:  p.GetProfile(),
			Ping:     p.GetPing(),
			Name:     chat.New(p.GetProfile().Name),
			Gamemode: player.CREATIVE,
		},
	}}
	conn.SendPacket(&addPlayer)
}

func debugPackets(p ents.Player, conn network.Connection) {
	conn.SendPacket(&play.PacketPlayOutTabInfo{
		Header: chat.Message{
			Text: "§b§lGo Server",
		},
		Footer: chat.Message{
			Text: "§f1.8",
		},
	})

	p.SendMsgColor(
		" ",
		" &b&lGo Server &f- &71.8",
		" ",
		" &fFollow the project on github:",
		" &bhttps://github.com/MineLC/Go-Server",
	)

	p.SendMsgColorPos(chat.HotBarText, "&b&lGo Server")
	p.SetXP(25)

	// See a example of sidebar in: https://github.com/ichocomilk/LightSidebar/blob/main/src/main/java/io/github/ichocomilk/lightsidebar/nms/v1_8R3/Sidebar1_8R3.java
	create := play.PacketPlayOutScoreboardObjective{
		Objective:            "sidebar",
		ObjectiveDisplayName: "§b§lGo server",
		Display:              play.INTEGER,
		Id:                   play.CREATE,
	}
	display := play.PacketPlayOutScoreboardDisplayObjective{
		Objective: "sidebar",
		Id:        play.SIDEBAR,
	}
	line := play.PacketPlayOutScoreboardScore{
		Line:      "Disable: §bconfig.toml",
		Objective: "sidebar",
		Score:     1,
		Remove:    false,
	}
	conn.SendPacket(&create)
	conn.SendPacket(&display)
	conn.SendPacket(&line)

	conn.SendPacket(&play.PacketPlayOutHeldItemChange{
		Slot: player.SLOT_0,
	})
	conn.SendPacket(&play.PacketPlayOutPosition{
		Position: data.PositionF{
			X: 0,
			Y: 50,
			Z: 0,
		},
		Head: data.HeadPosition{
			AxisX: 0,
			AxisY: 0,
		},
	})
}
