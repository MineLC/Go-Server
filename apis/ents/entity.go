package ents

import "github.com/golangmc/minecraft-server/apis/base"

type Entity interface {
	base.Unique

	EntityUUID() int32
}
