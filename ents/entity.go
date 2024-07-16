package ents

import "github.com/minelc/go-server-api/data"

var entityCounter = int64(0)

type entity struct {
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
