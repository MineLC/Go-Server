package server

import (
	"github.com/minelc/go-server/api/data/player"
	"github.com/minelc/go-server/api/game"
	"github.com/minelc/go-server/api/network"
	"github.com/minelc/go-server/api/world"
)

type PacketPlayOutLogin struct {
	EntityID    int32
	Hardcore    bool
	GameMode    player.GameMode
	Dimension   world.Dimension
	Difficulty  game.Difficulty
	MaxPlayers  byte
	LevelType   world.LevelType
	ReduceDebug bool
}

func (p *PacketPlayOutLogin) UUID() int32 {
	return 1
}

func (p *PacketPlayOutLogin) Push(writer network.Buffer) {
	writer.PushI32(p.EntityID)

	gamemode := byte(p.GameMode)
	if p.Hardcore {
		gamemode |= 0x8
	}

	writer.PushByt(gamemode)
	writer.PushByt(byte(p.Dimension))
	writer.PushByt(byte(p.Difficulty))
	writer.PushByt(p.MaxPlayers)
	writer.PushTxt(p.LevelType.String())
	writer.PushBit(p.ReduceDebug)
}
