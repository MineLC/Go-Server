package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/apis/data"
	"github.com/golangmc/minecraft-server/impl/base"
	"github.com/golangmc/minecraft-server/impl/data/client"
)

type PacketPlayOutPosition struct {
	Location data.Location
	Relative client.Relativity

	//ID int32
}

func (p *PacketPlayOutPosition) UUID() int32 {
	return 8
}

func (p *PacketPlayOutPosition) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushF64(p.Location.X)
	writer.PushF64(p.Location.Y)
	writer.PushF64(p.Location.Z)

	writer.PushF32(p.Location.AxisX)
	writer.PushF32(p.Location.AxisY)

	p.Relative.Push(writer)

	//writer.PushVrI(p.ID)
}
