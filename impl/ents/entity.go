package ents

var entityCounter = int64(0)

type entity struct {
	name     string
	entityID int64
}

func NewEntity() entity {
	id := entityCounter
	entityCounter++

	return entity{entityID: id}
}

func (e *entity) EntityUUID() int64 {
	return e.entityID
}
