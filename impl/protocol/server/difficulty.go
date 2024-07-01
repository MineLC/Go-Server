package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/apis/game"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketPlayOutServerDifficulty struct {
	Difficulty game.Difficulty
}

func (p *PacketPlayOutServerDifficulty) UUID() int32 {
	return 65
}

func (p *PacketPlayOutServerDifficulty) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushByt(byte(p.Difficulty))
}
