package player

import "github.com/minelc/go-server/api/data"

type Profile struct {
	UUID data.UUID
	Name string

	Properties []*ProfileProperty
}

type ProfileProperty struct {
	Name      string
	Value     string
	Signature *string
}
