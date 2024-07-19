package ents

import (
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/network"
)

type entityLiving struct {
	entity

	nametag *string
	health  float64

	head data.HeadPosition
}

func NewEntityLiving() entityLiving {
	return entityLiving{entity: NewEntity()}
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

func (e *entityLiving) PushMetadata(buffer network.Buffer) {
	e.entity.PushMetadata(buffer)

	buffer.PushByt(NameTag)
	if e.nametag == nil {
		buffer.PushTxt("")
		buffer.PushByt(ShowNameTag)
		buffer.PushByt(False)
	} else {
		buffer.PushTxt(*e.nametag)
		buffer.PushByt(ShowNameTag)
		buffer.PushByt(True)
	}

	buffer.PushByt(Health)
	buffer.PushF32(float32(e.health))

	buffer.PushByt(HasIA)
	buffer.PushByt(False)
}
