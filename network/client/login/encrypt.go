package login

import (
	"bytes"
	"fmt"

	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/data/chat"
	"github.com/minelc/go-server-api/data/player"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/client/login"
	srv_login "github.com/minelc/go-server-api/network/server/login"
	"github.com/minelc/go-server/ents"
	"github.com/minelc/go-server/game/join"
	"github.com/minelc/go-server/network/crypto/auth"
)

func HandleEncryption(conn network.Connection, packet network.PacketI) {
	p := packet.(*login.PacketIEncryptionResponse)
	defer func() {
		if err := recover(); err != nil {
			conn.SendPacket(&srv_login.PacketODisconnect{
				Reason: *chat.New(fmt.Sprintf("Authentication failed: %v", err)).SetColor(chat.Red),
			})
		}
	}()

	ver, err := auth.Decrypt(p.Verify)
	if err != nil {
		panic(fmt.Errorf("failed to decrypt token: %s %v", conn.CertifyName(), err.Error()))
	}

	if !bytes.Equal(ver, conn.CertifyData()) {
		panic(fmt.Errorf("encryption failed, tokens are different: %s\n%v | %v", conn.CertifyName(), ver, conn.CertifyData()))
	}

	sec, err := auth.Decrypt(p.Secret)
	if err != nil {
		panic(fmt.Errorf("failed to decrypt secret: %s %v", conn.CertifyName(), err.Error()))
	}

	conn.CertifyUpdate(sec) // enable encryption on the connection
	auth.RunAuthGet(sec, conn.CertifyName(), func(auth *auth.Auth, err error) {
		defer func() {
			if err := recover(); err != nil {
				conn.SendPacket(&srv_login.PacketODisconnect{
					Reason: *chat.New(fmt.Sprintf("Authentication failed: %v", err)).SetColor(chat.Red),
				})
			}
		}()

		if err != nil {
			panic(fmt.Errorf("failed to authenticate: %s - %v", conn.CertifyName(), err))
		}

		uuid, err := data.TextToUUID(auth.UUID)
		if err != nil {
			panic(fmt.Errorf("failed to decode uuid for %s: %s - %v", conn.CertifyName(), auth.UUID, err.Error()))
		}

		prof := player.Profile{
			UUID: uuid,
			Name: auth.Name,
		}

		for _, prop := range auth.Prop {
			prof.Properties = append(prof.Properties, &player.ProfileProperty{
				Name:      prop.Name,
				Value:     prop.Data,
				Signature: prop.Sign,
			})
		}
		join.Join(ents.NewPlayer(&prof, conn), conn)
	})
}
