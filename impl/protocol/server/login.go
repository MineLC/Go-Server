package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/apis/game"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketPlayOutLogin struct {
	EntityID    int32
	Hardcore    bool
	GameMode    game.GameMode
	Dimension   game.Dimension
	Difficulty  game.Difficulty
	MaxPlayers  byte
	LevelType   game.LevelType
	ReduceDebug bool
}

func (p *PacketPlayOutLogin) UUID() int32 {
	return 1
}

func (p *PacketPlayOutLogin) Push(writer buff.Buffer, conn base.Connection) {
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
