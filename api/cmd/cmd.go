package cmd

import "github.com/minelc/go-server/api/ents"

type Command interface {
	Execute(sender ents.Sender, args []string)
}

type StructCommand struct {
	Execute func(sender ents.Sender, args []string)
	Tab     func(sender ents.Sender, args []string) []string
}

type TabCommand interface {
	Command
	Tab(sender ents.Sender, args []string) []string
}
