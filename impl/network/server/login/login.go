package login

import (
	"github.com/minelc/go-server/api/network"
)

type PacketOEncryptionRequest struct {
	Server string // unused?
	Public []byte
	Verify []byte
}

func (p *PacketOEncryptionRequest) UUID() int32 {
	return 0x01
}

func (p *PacketOEncryptionRequest) Push(writer network.Buffer) {
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

func (p *PacketOLoginSuccess) Push(writer network.Buffer) {
	writer.PushTxt(p.PlayerUUID)
	writer.PushTxt(p.PlayerName)
}

type PacketOSetCompression struct {
	Threshold int32
}

func (p *PacketOSetCompression) UUID() int32 {
	return 0x03
}

func (p *PacketOSetCompression) Push(writer network.Buffer) {
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

func (p *PacketOLoginPluginRequest) Push(writer network.Buffer, conn network.Connection) {
	writer.PushVrI(p.MessageID)
	writer.PushTxt(p.Channel)
	writer.PushUAS(p.OptData, false)
}
