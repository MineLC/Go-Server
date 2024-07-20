package login

import (
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/data/player"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/client/login"
	srv_log "github.com/minelc/go-server-api/network/server/login"
	"github.com/minelc/go-server/ents"
	"github.com/minelc/go-server/game/join"
	"github.com/minelc/go-server/network/crypto/auth"
	uuid "github.com/satori/go.uuid"
)

func HandleLoginStart(conn network.Connection, packet network.PacketI) {
	p := packet.(*login.PacketILoginStart)
	conn.CertifyValues(p.PlayerName)

	if join.GameOptions.OnlineMode {
		_, public := auth.NewCrypt()

		response := srv_log.PacketOEncryptionRequest{
			Server: "",
			Public: public,
			Verify: conn.CertifyData(),
		}

		conn.SendPacket(&response)
		return
	}
	prof := player.Profile{
		UUID: uuid.FromStringOrNil(string(data.CreateUUID(p.PlayerName))),
		Name: p.PlayerName,
	}

	join.Join(ents.NewPlayer(&prof, conn), conn)
}
