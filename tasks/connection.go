package tasks

import (
	"time"

	"github.com/minelc/go-server-api/ents"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/server/play"
)

var lastKeepAlive *int64

func KeepAlive(p *map[network.Connection]ents.Player) error {
	players := *p
	now := time.Now().Unix()

	packet := &play.PacketPlayOutKeepAlive{KeepAliveID: int32(now / 1e6)}

	for conn, p := range players {
		if p.GetKeepAlive() != -1 {
			if now-p.GetKeepAlive() >= 20000 { // In millis
				p.Disconnect()
				continue
			}
		}
		conn.SendPacket(packet)
	}
	lastKeepAlive = &now
	return nil
}

func GetLastAlive() int64 {
	return *lastKeepAlive
}
