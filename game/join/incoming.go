package join

import (
	"github.com/minelc/go-server-api/data"
	entities "github.com/minelc/go-server-api/data/entities"
	"github.com/minelc/go-server-api/ents"
	"github.com/minelc/go-server-api/network"
)

/*
This packets will be add in the api.
Now, this are for testing and debug
*/
type PacketPlayOutEntityMetadata struct {
	Entity ents.Entity
}

func (p *PacketPlayOutEntityMetadata) UUID() int32 {
	return 0x1C
}

func (p *PacketPlayOutEntityMetadata) Push(writer network.Buffer) {
	writer.PushVrI(int32(p.Entity.EntityUUID()))
	p.Entity.PushMetadata(writer)
	writer.PushByt(127)
}

type PacketPlayOutSpawnEntityLiving struct {
	Entity ents.EntityLiving
	Type   entities.CREATURE
}

func (p *PacketPlayOutSpawnEntityLiving) UUID() int32 {
	return 15
}

func (p *PacketPlayOutSpawnEntityLiving) Push(writer network.Buffer) {
	writer.PushVrI(int32(p.Entity.EntityUUID()))
	writer.PushByt(byte(p.Type & 255))

	writer.PushI32(int32(p.Entity.GetPosition().X * 32.0))
	writer.PushI32(int32(p.Entity.GetPosition().Y * 32.0))
	writer.PushI32(int32(p.Entity.GetPosition().Z * 32.0))

	writer.PushByt(byte(p.Entity.GetHeadPos().AxisX * 256.0 / 360.0))
	writer.PushByt(byte(p.Entity.GetHeadPos().AxisY * 256.0 / 360.0))
	writer.PushByt(byte(p.Entity.GetHeadPos().AxisX * 256.0 / 360.0))

	writer.PushI16(0)
	writer.PushI16(0)
	writer.PushI16(0)

	p.Entity.PushMetadata(writer)
	writer.PushByt(127)
}

type PacketPlayOutSpawnPlayer struct {
	Player ents.Player
}

func (p *PacketPlayOutSpawnPlayer) UUID() int32 {
	return 0x0C
}

func (p *PacketPlayOutSpawnPlayer) Push(writer network.Buffer) {
	pos := p.Player.GetPosition()

	writer.PushVrI(int32(p.Player.EntityUUID()))
	writer.PushUID(p.Player.GetProfile().UUID)
	writer.PushPos(data.PositionI{
		X: int64(pos.X),
		Y: int64(pos.Y),
		Z: int64(pos.Z),
	})
	writer.PushByt(0)
	writer.PushByt(0)

	writer.PushI16(0)

	p.Player.PushMetadata(writer)
	writer.PushByt(127)
}

type PacketPlayOutExperience struct {
	Xp    int32
	Level int32
}

func (p *PacketPlayOutExperience) UUID() int32 {
	return 0x1F
}

func (p *PacketPlayOutExperience) Push(writer network.Buffer) {
	nextLevel := GetXP(p.Level + 1)
	if p.Xp == 0 {
		writer.PushF32(0.0)
	} else {
		writer.PushF32(float32(p.Xp) / float32(nextLevel))
	}
	writer.PushVrI(p.Level)
	writer.PushVrI(p.Xp)
}

func GetXP(level int32) int32 {
	if level <= 16 {
		return level*level + 6*level
	} else if level <= 31 {
		return int32(2.5*float32(level*level)) + int32(40.5*float32(level)) + 360
	} else {
		return int32(4.5*float32(level*level)) + int32(162.5*float32(level)) + 2220
	}
}

func GetLevel(xp int) int32 {
	i := 0

	for {
		if xp < 0 {
			break
		}
		if i < 16 {
			xp -= (2 * i) + 7
		} else if i < 31 {
			xp -= (5 * i) - 38
		} else {
			xp -= (9 * i) - 158
		}
		i++
	}
	return int32(i - 1)
}
