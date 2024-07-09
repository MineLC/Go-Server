package server

import (
	"github.com/minelc/go-server/api/game"
	"github.com/minelc/go-server/api/network"
)

type PacketPlayOutServerDifficulty struct {
	Difficulty game.Difficulty
}

func (p *PacketPlayOutServerDifficulty) UUID() int32 {
	return 65
}

func (p *PacketPlayOutServerDifficulty) Push(writer network.Buffer) {
	writer.PushByt(byte(p.Difficulty))
}
