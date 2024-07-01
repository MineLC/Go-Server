package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketPlayOutUpdateHealth struct {
	ScaledHealth float32
	FoodLevel    int32
	Saturation   float32
}

func (p *PacketPlayOutUpdateHealth) UUID() int32 {
	return 6
}

func (p *PacketPlayOutUpdateHealth) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushF32(p.ScaledHealth)
	writer.PushVrI(p.FoodLevel)
	writer.PushF32(p.Saturation)
}
