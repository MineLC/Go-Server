package play

import (
	"github.com/minelc/go-server/api/network"
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

func (p *PacketPlayInFlying) Pull(reader network.Buffer) {
	p.OnGround = reader.PullBit()
}

func (p *PacketPlayInFlying) Handle(c *network.Connection) {

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

func (p *PacketPlayInLook) Pull(reader network.Buffer) {
	p.Yaw = reader.PullF32()
	p.Pitch = reader.PullF32()
	p.OnGround = reader.PullBit()
}

func (p *PacketPlayInLook) Handle(c *network.Connection) {

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

func (p *PacketPlayInPosition) Pull(reader network.Buffer) {
	p.X = reader.PullF64()
	p.Y = reader.PullF64()
	p.Z = reader.PullF64()
	p.OnGround = reader.PullBit()
}

func (p *PacketPlayInPosition) Handle(c *network.Connection) {

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

func (p *PacketPlayInPositionLook) Pull(reader network.Buffer) {
	p.X = reader.PullF64()
	p.Y = reader.PullF64()
	p.Z = reader.PullF64()
	p.Yaw = reader.PullF32()
	p.Pitch = reader.PullF32()
	p.OnGround = reader.PullBit()
}

func (p *PacketPlayInPositionLook) Handle(c *network.Connection) {

}
