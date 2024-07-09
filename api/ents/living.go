package ents

import "github.com/minelc/go-server/api/data"

type EntityLiving interface {
	Entity

	GetHealth() float64
	SetHealth(health float64)
	UUID() data.UUID
}
