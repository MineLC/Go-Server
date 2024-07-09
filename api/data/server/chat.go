package server

type ChatMode int

const (
	Full ChatMode = iota
	Cmds
	None
)
