package login

import (
	"bytes"
	"encoding/hex"

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
		UUID: uuid.FromStringOrNil(string(createUUID(p.PlayerName))),
		Name: p.PlayerName,
	}

	join.Join(ents.NewPlayer(&prof, conn), conn)
}

func createUUID(playerName string) []byte {
	var buffer bytes.Buffer

	nameInHex := hex.EncodeToString([]byte(playerName))
	length := len(nameInHex)

	if length <= 12 {
		buffer.WriteString("00000000-0000-4000-0000-")
		buffer.WriteString(nameInHex)
		diference := 12 - length
		for i := 0; i < diference; i++ {
			buffer.WriteRune('0')
		}
		return buffer.Bytes()
	}
	if length <= 20 {
		buffer.WriteString("00000000-4000-")
		buffer.WriteString(nameInHex[:4])
		buffer.WriteRune('-')
		buffer.WriteString(nameInHex[4:8])
		buffer.WriteRune('-')
		buffer.WriteString(nameInHex[8:])

		diference := 36 - buffer.Len()
		for i := 0; i < diference; i++ {
			buffer.WriteRune('0')
		}
		return buffer.Bytes()
	}
	buffer.WriteString(nameInHex[:8])
	buffer.WriteRune('-')
	buffer.WriteString(nameInHex[8:12])
	buffer.WriteRune('-')
	buffer.WriteString(nameInHex[12:16])
	buffer.WriteRune('-')
	buffer.WriteString(nameInHex[16:20])
	buffer.WriteRune('-')
	buffer.WriteString(nameInHex[20:])

	diference := 36 - buffer.Len()
	for i := 0; i < diference; i++ {
		buffer.WriteRune('0')
	}
	return buffer.Bytes()
}
