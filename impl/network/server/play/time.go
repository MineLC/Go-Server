package server

import "github.com/minelc/go-server/api/network"

type PacketPlayOutUpdateTime struct {
	Time    int64
	DayTime int64
}

func (p *PacketPlayOutUpdateTime) UUID() int32 {
	return 3
}

func (p *PacketPlayOutUpdateTime) Push(writer network.Buffer) {
	writer.PushI64(p.Time)
	writer.PushI64(p.DayTime)
}
