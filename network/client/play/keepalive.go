package play

import (
	"time"

	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server/tasks"
)

func HandleKeepAlive(c *network.Connection, _ network.PacketI) {
	player := api.GetServer().GetPlayer(*c)
	if player != nil {
		now := time.Now().Unix()
		(*player).SetKeepAlive(now)
		(*player).SetPing(now, tasks.GetLastAlive())
	}
}
