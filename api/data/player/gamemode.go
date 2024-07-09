package player

type GameMode byte

const (
	SURVIVAL GameMode = iota
	CREATIVE
	ADVENTURE
	SPECTATOR
)
