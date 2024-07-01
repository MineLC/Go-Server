package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

// done

type PacketLoginInStart struct {
	PlayerName string
}

func (p *PacketLoginInStart) UUID() int32 {
	return 0
}

func (p *PacketLoginInStart) Pull(reader buff.Buffer, conn base.Connection) {
	p.PlayerName = reader.PullTxt()
}

type PacketLoginInEncryptionBegin struct {
	Secret []byte
	Verify []byte
}

func (p *PacketLoginInEncryptionBegin) UUID() int32 {
	return 1
}

func (p *PacketLoginInEncryptionBegin) Pull(reader buff.Buffer, conn base.Connection) {
	p.Secret = reader.PullUAS()
	p.Verify = reader.PullUAS()
}

type PacketILoginPluginResponse struct {
	Message int32
	Success bool
	OptData []byte
}

func (p *PacketILoginPluginResponse) UUID() int32 {
	return 0x02
}

func (p *PacketILoginPluginResponse) Pull(reader buff.Buffer, conn base.Connection) {
	p.Message = reader.PullVrI()
	p.Success = reader.PullBit()
	p.OptData = reader.UAS()[reader.InI():reader.Len()]
}
