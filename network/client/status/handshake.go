package status

import (
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/data/chat"
	"github.com/minelc/go-server-api/network"
	cli_status "github.com/minelc/go-server-api/network/client/status"
	"github.com/minelc/go-server-api/network/server/login"
	"github.com/minelc/go-server-api/network/server/status"
)

func HandleHandShake(c *network.Connection, packet network.PacketI) {
	p := packet.(*cli_status.PacketHandshakingInSetProtocol)
	conn := (*c)

	if p.State == network.STATUS {
		conn.SetState(network.STATUS)
		conn.SendPacket(&status.PacketOResponse{Motd: data.Server.Motd})
		return
	}
	if p.State != network.LOGIN {
		conn.Stop()
		return
	}

	if p.Version > 47 {
		conn.SendPacket(&login.PacketODisconnect{Reason: *chat.New("New client. Use: 1.8")})
		conn.Stop()
		return
	}
	if p.Version < 47 {
		conn.SendPacket(&login.PacketODisconnect{Reason: *chat.New("Old client. Use: 1.8")})
		conn.Stop()
		return
	}

	conn.SetState(network.LOGIN)
}
