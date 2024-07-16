package login

import (
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/client/login"
	srv_log "github.com/minelc/go-server-api/network/server/login"
	"github.com/minelc/go-server/network/crypto/auth"
)

func HandleLoginStart(conn network.Connection, packet network.PacketI) {
	p := packet.(*login.PacketILoginStart)
	conn.CertifyValues(p.PlayerName)

	_, public := auth.NewCrypt()

	response := srv_log.PacketOEncryptionRequest{
		Server: "",
		Public: public,
		Verify: conn.CertifyData(),
	}

	conn.SendPacket(&response)
}
