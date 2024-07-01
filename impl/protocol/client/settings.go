package client

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

type SETTING byte

const (
	FULL SETTING = iota
	SYSTEM
	HIDDEN
)

type PacketPlayInSettings struct {
	A string
	b byte
	c SETTING
	d bool
	e uint8
}

func (p *PacketPlayInSettings) UUID() int32 {
	return 21
}

func (p *PacketPlayInSettings) Pull(reader buff.Buffer, conn base.Connection) {
	p.A = reader.PullTxt()
	p.b = reader.PullByt()
	p.c = SETTING(reader.PullByt())
	p.d = reader.PullBit()
	p.e = reader.PullByt()
}
