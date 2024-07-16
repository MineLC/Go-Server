package ents

import "github.com/minelc/go-server-api/data"

type entityLiving struct {
	entity

	health float64
	head   data.HeadPosition
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
