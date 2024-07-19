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

func (e *entity) PushMetadata(buffer network.Buffer) {
	buffer.PushByt(BitMask)
	buffer.PushByt(255)
}
