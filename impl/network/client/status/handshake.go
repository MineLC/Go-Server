package status

import (
	"github.com/minelc/go-server/api/data/chat"
	"github.com/minelc/go-server/api/data/motd"
	"github.com/minelc/go-server/api/network"
	"github.com/minelc/go-server/impl/network/server/login"
	"github.com/minelc/go-server/impl/network/server/status"
)

type PacketHandshakingInSetProtocol struct {
	Version int32

	host string
	port uint16

	State network.PacketState
}

func (p *PacketHandshakingInSetProtocol) Pull(reader network.Buffer) {
	p.Version = reader.PullVrI()

	p.host = reader.PullTxt()
	p.port = reader.PullU16()

	p.State = network.StateValueOf(int(reader.PullVrI()))
}

func (p *PacketHandshakingInSetProtocol) UUID() int32 {
	return 0
}

func (p *PacketHandshakingInSetProtocol) Handle(c *network.Connection) {
	conn := (*c)
	if p.State == network.STATUS {
		conn.SetState(network.STATUS)
		conn.SendPacket(&status.PacketOResponse{Motd: motd.GetResponse()})
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
