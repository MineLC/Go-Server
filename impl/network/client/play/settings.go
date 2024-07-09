package play

import (
	"github.com/minelc/go-server/api/network"
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

func (p *PacketPlayInSettings) Pull(reader network.Buffer) {
	p.A = reader.PullTxt()
	p.b = reader.PullByt()
	p.c = SETTING(reader.PullByt())
	p.d = reader.PullBit()
	p.e = reader.PullByt()
}

func (p *PacketPlayInSettings) Handle(c *network.Connection) {

}
