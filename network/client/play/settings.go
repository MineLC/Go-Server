package play

import (
	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/client/play"
	"github.com/minelc/go-server/game/join"
)

func HandleSettings(c network.Connection, packet network.PacketI) {
	p := packet.(*play.PacketPlayInSettings)
	player := api.GetServer().GetPlayer(c)
	settings := player.GetClientSettings()

	if p.SkinBitMask == settings.SkinParts &&
		p.ChatMode == play.ChatMode(settings.ChatMode) &&
		p.Language == settings.Language &&
		p.ViewDistance == settings.ViewDistance {
		return
	}

	settings.ViewDistance = p.ViewDistance
	settings.ChatMode = byte(p.ChatMode)
	settings.Language = p.Language

	if p.SkinBitMask != settings.SkinParts {
		settings.SkinParts = p.SkinBitMask
		c.SendPacket(&join.PacketPlayOutEntityMetadata{Entity: player})
	}
}
