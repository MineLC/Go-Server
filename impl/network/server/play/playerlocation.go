package server

import (
	"github.com/minelc/go-server/api/data"
	"github.com/minelc/go-server/api/network"
)

type PacketPlayOutPosition struct {
	Location data.Location
}

func (p *PacketPlayOutPosition) UUID() int32 {
	return 8
}

func (p *PacketPlayOutPosition) Push(writer network.Buffer) {
	writer.PushF64(p.Location.X)
	writer.PushF64(p.Location.Y)
	writer.PushF64(p.Location.Z)

	writer.PushF32(p.Location.AxisX)
	writer.PushF32(p.Location.AxisY)

	// No relativity
	writer.PushByt(0)
}
