package debug

import (
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/data/chat"
	"github.com/minelc/go-server-api/data/player"
	"github.com/minelc/go-server-api/ents"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/server/play"
)

func SendDebugPackets(p ents.Player, conn network.Connection) {
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

	conn.SendPacket(&play.PacketPlayOutHeldItemChange{
		Slot: player.SLOT_0,
	})
	conn.SendPacket(&play.PacketPlayOutPosition{
		Location: data.Location{
			X:     0,
			Y:     10,
			Z:     0,
			AxisX: 0,
			AxisY: 0,
		},
	})
}
