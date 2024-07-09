package ents

import (
	"github.com/minelc/go-server/api/data/player"
)

type Player interface {
	EntityLiving
	Sender

	GetIsOnline() bool
	SetIsOnline(state bool)

	GetGamemode() player.GameMode
	SetGamemode(gamemode player.GameMode)

	GetProfile() *player.Profile

	GetPing() int32
	SetPing(ping_delay int64, server_ping int64)
	GetKeepAlive() int64
	SetKeepAlive(time int64)

	Disconnect()
}
