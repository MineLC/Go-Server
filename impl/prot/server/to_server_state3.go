package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/apis/data"
	"github.com/golangmc/minecraft-server/apis/game"
	"github.com/golangmc/minecraft-server/impl/base"
	"github.com/golangmc/minecraft-server/impl/data/client"
	"github.com/golangmc/minecraft-server/impl/data/plugin"
)

type PacketIKeepAlive struct {
	KeepAliveID int64
}

func (p *PacketIKeepAlive) UUID() int32 {
	return 0
}

func (p *PacketIKeepAlive) Pull(reader buff.Buffer, conn base.Connection) {
	p.KeepAliveID = reader.PullI64()
}

type PacketIChatMessage struct {
	Message string
}

func (p *PacketIChatMessage) UUID() int32 {
	return 1
}

func (p *PacketIChatMessage) Pull(reader buff.Buffer, conn base.Connection) {
	p.Message = reader.PullTxt()
}

type PacketITeleportConfirm struct {
	TeleportID int32
}

func (p *PacketITeleportConfirm) UUID() int32 {
	return 0x00
}

func (p *PacketITeleportConfirm) Pull(reader buff.Buffer, conn base.Connection) {
	p.TeleportID = reader.PullVrI()
}

type PacketIQueryBlockNBT struct {
	TransactionID int32
	Position      data.PositionI
}

func (p *PacketIQueryBlockNBT) UUID() int32 {
	return 0x01
}

func (p *PacketIQueryBlockNBT) Pull(reader buff.Buffer, conn base.Connection) {
	p.TransactionID = reader.PullVrI()
	p.Position = reader.PullPos()
}

type PacketISetDifficulty struct {
	Difficult game.Difficulty
}

func (p *PacketISetDifficulty) UUID() int32 {
	return 2
}

func (p *PacketISetDifficulty) Pull(reader buff.Buffer, conn base.Connection) {
	p.Difficult = game.DifficultyValueOf(reader.PullByt())
}

type PacketIPluginMessage struct {
	Message plugin.Message
}

func (p *PacketIPluginMessage) UUID() int32 {
	return 0x0B
}

func (p *PacketIPluginMessage) Pull(reader buff.Buffer, conn base.Connection) {
	channel := reader.PullTxt()
	message := plugin.GetMessageForChannel(channel)

	if message == nil {
		return // log unregistered channel?
	}

	message.Pull(reader)

	p.Message = message
}

type PacketIClientStatus struct {
	Action client.StatusAction
}

func (p *PacketIClientStatus) UUID() int32 {
	return 0x04
}

func (p *PacketIClientStatus) Pull(reader buff.Buffer, conn base.Connection) {
	p.Action = client.StatusAction(reader.PullVrI())
}

type PacketIClientSettings struct {
	Locale       string
	ViewDistance byte
	ChatMode     client.ChatMode
	ChatColors   bool // if false, strip messages of colors before sending
	SkinParts    client.SkinParts
	MainHand     client.MainHand
}

func (p *PacketIClientSettings) UUID() int32 {
	return 0x05
}

func (p *PacketIClientSettings) Pull(reader buff.Buffer, conn base.Connection) {
	p.Locale = reader.PullTxt()
	p.ViewDistance = reader.PullByt()
	p.ChatMode = client.ChatMode(reader.PullVrI())
	p.ChatColors = reader.PullBit()

	parts := client.SkinParts{}
	parts.Pull(reader)

	p.SkinParts = parts
	p.MainHand = client.MainHand(reader.PullVrI())
}

type PacketIPlayerPosition struct {
	Position data.PositionF
	OnGround bool
}

func (p *PacketIPlayerPosition) UUID() int32 {
	return 0x11
}

func (p *PacketIPlayerPosition) Pull(reader buff.Buffer, conn base.Connection) {
	p.Position = data.PositionF{
		X: reader.PullF64(),
		Y: reader.PullF64(),
		Z: reader.PullF64(),
	}

	p.OnGround = reader.PullBit()
}

type PacketIPlayerLocation struct {
	Location data.Location
	OnGround bool
}

func (p *PacketIPlayerLocation) UUID() int32 {
	return 0x12
}

func (p *PacketIPlayerLocation) Pull(reader buff.Buffer, conn base.Connection) {
	p.Location = data.Location{
		PositionF: data.PositionF{
			X: reader.PullF64(),
			Y: reader.PullF64(),
			Z: reader.PullF64(),
		},
		RotationF: data.RotationF{
			AxisX: reader.PullF32(),
			AxisY: reader.PullF32(),
		},
	}

	p.OnGround = reader.PullBit()
}

type PacketIPlayerRotation struct {
	Rotation data.RotationF
	OnGround bool
}

func (p *PacketIPlayerRotation) UUID() int32 {
	return 0x13
}

func (p *PacketIPlayerRotation) Pull(reader buff.Buffer, conn base.Connection) {
	p.Rotation = data.RotationF{
		AxisX: reader.PullF32(),
		AxisY: reader.PullF32(),
	}

	p.OnGround = reader.PullBit()
}
