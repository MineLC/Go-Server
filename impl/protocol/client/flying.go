package client

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/impl/base"
)

/*
Send if player is flying
*/
type PacketPlayInFlying struct {
	OnGround bool
}

func (p *PacketPlayInFlying) UUID() int32 {
	return 3
}

func (p *PacketPlayInFlying) Pull(reader buff.Buffer, conn base.Connection) {
	p.OnGround = reader.PullBit()
}

/*
Send only if the player move the camera
*/
type PacketPlayInLook struct {
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (p *PacketPlayInLook) UUID() int32 {
	return 5
}

func (p *PacketPlayInLook) Pull(reader buff.Buffer, conn base.Connection) {
	p.Yaw = reader.PullF32()
	p.Pitch = reader.PullF32()
	p.OnGround = reader.PullBit()
}

/*
Send only if the player move around x,y or z cords
*/
type PacketPlayInPosition struct {
	X        float64
	Y        float64
	Z        float64
	OnGround bool
}

func (p *PacketPlayInPosition) UUID() int32 {
	return 4
}

func (p *PacketPlayInPosition) Pull(reader buff.Buffer, conn base.Connection) {
	p.X = reader.PullF64()
	p.Y = reader.PullF64()
	p.Z = reader.PullF64()
	p.OnGround = reader.PullBit()
}

/*
Send if the player move around x,y or z cords and change the camera
*/
type PacketPlayInPositionLook struct {
	X        float64
	Y        float64
	Z        float64
	Yaw      float32
	Pitch    float32
	OnGround bool
}

func (p *PacketPlayInPositionLook) UUID() int32 {
	return 6
}

func (p *PacketPlayInPositionLook) Pull(reader buff.Buffer, conn base.Connection) {
	p.X = reader.PullF64()
	p.Y = reader.PullF64()
	p.Z = reader.PullF64()
	p.Yaw = reader.PullF32()
	p.Pitch = reader.PullF32()
	p.OnGround = reader.PullBit()
}
