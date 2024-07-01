package client

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/apis/data/msgs"
	"github.com/golangmc/minecraft-server/impl/base"
)

// done

type PacketODisconnect struct {
	Reason msgs.Message
}

func (p *PacketODisconnect) UUID() int32 {
	return 0x00
}

func (p *PacketODisconnect) Push(writer buff.Buffer, conn base.Connection) {
	message := p.Reason

	writer.PushTxt(message.AsJson())
}

type PacketOEncryptionRequest struct {
	Server string // unused?
	Public []byte
	Verify []byte
}

func (p *PacketOEncryptionRequest) UUID() int32 {
	return 0x01
}

func (p *PacketOEncryptionRequest) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushTxt(p.Server)
	writer.PushUAS(p.Public, true)
	writer.PushUAS(p.Verify, true)
}

type PacketOLoginSuccess struct {
	PlayerUUID string
	PlayerName string
}

func (p *PacketOLoginSuccess) UUID() int32 {
	return 0x02
}

func (p *PacketOLoginSuccess) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushTxt(p.PlayerUUID)
	writer.PushTxt(p.PlayerName)
}

type PacketOSetCompression struct {
	Threshold int32
}

func (p *PacketOSetCompression) UUID() int32 {
	return 0x03
}

func (p *PacketOSetCompression) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushVrI(p.Threshold)
}

type PacketOLoginPluginRequest struct {
	MessageID int32
	Channel   string
	OptData   []byte
}

func (p *PacketOLoginPluginRequest) UUID() int32 {
	return 0x04
}

func (p *PacketOLoginPluginRequest) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushVrI(p.MessageID)
	writer.PushTxt(p.Channel)
	writer.PushUAS(p.OptData, false)
}
