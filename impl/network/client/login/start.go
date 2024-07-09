package login

import (
	"github.com/minelc/go-server/api/network"
	"github.com/minelc/go-server/impl/network/crypto/auth"
	"github.com/minelc/go-server/impl/network/server/login"
)

type PacketILoginStart struct {
	PlayerName string
}

func (p *PacketILoginStart) UUID() int32 {
	return 0x00
}

func (p *PacketILoginStart) Pull(reader network.Buffer) {
	p.PlayerName = reader.PullTxt()
}

func (p *PacketILoginStart) Handle(conn *network.Connection) {
	(*conn).CertifyValues(p.PlayerName)

	_, public := auth.NewCrypt()

	response := login.PacketOEncryptionRequest{
		Server: "",
		Public: public,
		Verify: (*conn).CertifyData(),
	}

	(*conn).SendPacket(&response)
}
