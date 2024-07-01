package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

type PacketPlayOutUpdateTime struct {
	Time    int64
	DayTime int64
}

func (p *PacketPlayOutUpdateTime) UUID() int32 {
	return 3
}

func (p *PacketPlayOutUpdateTime) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushI64(p.Time)
	writer.PushI64(p.DayTime)
}
