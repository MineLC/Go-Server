package server

import "github.com/minelc/go-server/api/network"

type PacketPlayOutUpdateHealth struct {
	ScaledHealth float32
	FoodLevel    int32
	Saturation   float32
}

func (p *PacketPlayOutUpdateHealth) UUID() int32 {
	return 6
}

func (p *PacketPlayOutUpdateHealth) Push(writer network.Buffer) {
	writer.PushF32(p.ScaledHealth)
	writer.PushVrI(p.FoodLevel)
	writer.PushF32(p.Saturation)
}
