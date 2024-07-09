package game

import (
	"github.com/minelc/go-server/api/data"
	"github.com/minelc/go-server/api/data/player"
	"github.com/minelc/go-server/api/ents"
	"github.com/minelc/go-server/api/game"
	"github.com/minelc/go-server/api/network"
	"github.com/minelc/go-server/api/world"

	"github.com/minelc/go-server/impl/debug"
	server "github.com/minelc/go-server/impl/network/server/play"
)

func Join(p ents.Player, conn network.Connection) {
	conn.SendPacket(&server.PacketPlayOutLogin{
		EntityID:    int32(p.EntityUUID()),
		Hardcore:    false,
		GameMode:    player.CREATIVE,
		Dimension:   world.OVERWORLD,
		Difficulty:  game.EASY,
		MaxPlayers:  10,
		LevelType:   world.DEFAULT,
		ReduceDebug: false,
	})

	conn.SendPacket(&server.PacketPlayOutServerDifficulty{
		Difficulty: game.PEACEFUL,
	})

	debug.SendDebugPackets(p, conn)

	conn.SendPacket(&server.PacketPlayOutAbilities{
		Abilities: player.PlayerAbilities{
			Invulnerable: true,
			Flying:       true,
			AllowFlight:  true,
			InstantBuild: false,
		},
		FlyingSpeed: 0.05, // default value
		FieldOfView: 0.1,  // default value
	})

	conn.SendPacket(&server.PacketPlayOutHeldItemChange{
		Slot: player.SLOT_0,
	})
	conn.SendPacket(&server.PacketPlayOutPosition{
		Location: data.Location{
			X:     0,
			Y:     10,
			Z:     0,
			AxisX: 0,
			AxisY: 0,
		},
	})

	conn.SendPacket(&server.PacketPlayOutEntityMetadata{Entity: p})
}
