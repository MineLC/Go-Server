package ents

import (
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/network"
	uuid "github.com/satori/go.uuid"
)

type entityLiving struct {
	entity

	nametag *string
	health  float64

	head data.HeadPosition
}

func NewEntityLiving() entityLiving {
	var asa string = "asa"
	return entityLiving{entity: NewEntity(), nametag: &asa}
}

func (e *entityLiving) GetHealth() float64 {
	return e.health
}

func (e *entityLiving) SetHealth(health float64) {
	e.health = health
}
func (e *entityLiving) GetHeadPos() *data.HeadPosition {
	return &e.head
}

func (e *entityLiving) UUID() data.UUID {
	return uuid.Nil
}

func (e *entityLiving) PushMetadata(buffer network.Buffer) {
	e.entity.PushMetadata(buffer)

	if e.nametag != nil {
		addType(buffer, String, NameTag)
		buffer.PushTxt(*e.nametag)
	}
}
