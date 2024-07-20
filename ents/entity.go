package ents

import (
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/network"
)

var entityCounter = int64(0)

type entity struct {
	Bitmask  byte
	name     string
	entityID int64
	position data.PositionF
}

func NewEntity() entity {
	id := entityCounter
	entityCounter++

	return entity{entityID: id}
}

func (e *entity) EntityUUID() int64 {
	return e.entityID
}

func (e *entity) GetPosition() *data.PositionF {
	return &e.position
}

func (e *entity) SetData(data EntityMeta) {
	mask := &e.Bitmask
	set(mask, 0x01, data.OnFire)
	set(mask, 0x02, data.Crouched)
	set(mask, 0x08, data.Sprinting)
	set(mask, 0x10, data.Eating)
	set(mask, 0x20, data.Invisible)
}

func (e *entity) PushMetadata(buffer network.Buffer) {
	addType(buffer, Byte, 0)
	buffer.PushByt(e.Bitmask)
}

func addType(buffer network.Buffer, typeData MetadataType, index byte) {
	buffer.PushByt((byte(typeData)<<5 | index&31) & 255)
}

func set(mask *byte, field byte, when bool) {
	if when {
		*mask |= field
	}
}
