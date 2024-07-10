package status

import (
	"github.com/minelc/go-server-api/network"
	cli_status "github.com/minelc/go-server-api/network/client/status"
	"github.com/minelc/go-server-api/network/server/status"
)

func HandlePing(c *network.Connection, packet network.PacketI) {
	p := packet.(*cli_status.PacketIPing)
	(*c).SendPacket(&status.PacketStatusOPing{Ping: p.Ping})
}
