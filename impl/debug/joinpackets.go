package debug

import (
	"github.com/minelc/go-server/api/data/chat"
	"github.com/minelc/go-server/api/network"
	server "github.com/minelc/go-server/impl/network/server/play"
)

func SendDebugPackets(conn network.Connection) {
	conn.SendPacket(&server.PacketPlayOutTabInfo{
		Header: chat.Message{
			Text: "§b§lGo Server",
		},
		Footer: chat.Message{
			Text: "§f1.8",
		},
	})

	// See a example of sidebar in: https://github.com/ichocomilk/LightSidebar/blob/main/src/main/java/io/github/ichocomilk/lightsidebar/nms/v1_8R3/Sidebar1_8R3.java
	create := server.PacketPlayOutScoreboardObjective{
		Objective:            "sidebar",
		ObjectiveDisplayName: "§b§lGo server",
		Display:              server.INTEGER,
		Id:                   server.CREATE,
	}
	display := server.PacketPlayOutScoreboardDisplayObjective{
		Objective: "sidebar",
		Id:        server.SIDEBAR,
	}
	line := server.PacketPlayOutScoreboardScore{
		Line:      "Disable: §bconfig.toml",
		Objective: "sidebar",
		Score:     1,
		Remove:    false,
	}
	conn.SendPacket(&create)
	conn.SendPacket(&display)
	conn.SendPacket(&line)
}
