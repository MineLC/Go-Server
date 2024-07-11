package game

import (
	"github.com/minelc/go-server-api/data/player"
	"github.com/minelc/go-server-api/ents"
	"github.com/minelc/go-server-api/game"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/server/play"
	"github.com/minelc/go-server-api/world"

	"github.com/minelc/go-server/debug"
)

func Join(p ents.Player, conn network.Connection) {
	conn.SendPacket(&play.PacketPlayOutLogin{
		EntityID:    int32(p.EntityUUID()),
		Hardcore:    false,
		GameMode:    player.CREATIVE,
		Dimension:   world.OVERWORLD,
		Difficulty:  game.EASY,
		MaxPlayers:  10,
		LevelType:   world.DEFAULT,
		ReduceDebug: false,
	})

	conn.SendPacket(&play.PacketPlayOutServerDifficulty{
		Difficulty: game.NORMAL,
	})

	debug.SendDebugPackets(p, conn)

	conn.SendPacket(&play.PacketPlayOutEntityMetadata{Entity: p})
}
