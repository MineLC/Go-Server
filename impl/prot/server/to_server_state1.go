package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

// done

type PacketStatusInStart struct {
}

func (p *PacketStatusInStart) UUID() int32 {
	return 0
}

func (p *PacketStatusInStart) Pull(reader buff.Buffer, conn base.Connection) {
	// no fields
}

type PacketStatusInPing struct {
	Ping int64
}

func (p *PacketStatusInPing) UUID() int32 {
	return 1
}

func (p *PacketStatusInPing) Pull(reader buff.Buffer, conn base.Connection) {
	p.Ping = reader.PullI64()
}
