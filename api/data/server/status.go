package server

type StatusAction int

const (
	Respawn StatusAction = iota
	Request
)
