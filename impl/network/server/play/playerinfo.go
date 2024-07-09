package server

import (
	"github.com/minelc/go-server/api/data/chat"
	"github.com/minelc/go-server/api/data/player"
	"github.com/minelc/go-server/api/network"
)

type INFO_ACTION int32

const (
	ADD_PLAYER          INFO_ACTION = 1
	UPDATE_GAME_MODE    INFO_ACTION = 2
	UPDATE_LATENCY      INFO_ACTION = 3
	UPDATE_DISPLAY_NAME INFO_ACTION = 4
	REMOVE_PLAYER       INFO_ACTION = 5
)

type PlayerInfoData struct {
	Ping     int32
	Profile  *player.Profile
	Gamemode player.GameMode
	Name     *chat.Message
}

type PacketPlayOutPlayerInfo struct {
	Action  INFO_ACTION
	Players []PlayerInfoData
}

func (p *PacketPlayOutPlayerInfo) UUID() int32 {
	return 56
}

func (p *PacketPlayOutPlayerInfo) Push(writer network.Buffer) {
	if p.Players == nil {
		return
	}
	writer.PushVrI(int32(p.Action))
	writer.PushVrI(int32(len(p.Players)))

	for _, player := range p.Players {
		switch p.Action {

		case ADD_PLAYER:
			writer.PushUID(player.Profile.UUID)
			writer.PushTxt(player.Profile.Name)
			properties := player.Profile.Properties
			writer.PushVrI(int32(len(properties)))

			for _, property := range properties {
				writer.PushTxt(property.Name)
				writer.PushTxt(property.Value)
				if property.Signature != nil {
					writer.PushBit(true)
					writer.PushTxt(*property.Signature)
					continue
				}
				writer.PushBit(false)
			}

			writer.PushVrI(int32(player.Gamemode))
			writer.PushVrI(player.Ping)
			if player.Name == nil {
				writer.PushBit(false)
				continue
			}
			writer.PushBit(true)
			writer.PushTxt(player.Name.AsJson())
			continue

		case UPDATE_GAME_MODE:
			writer.PushUID(player.Profile.UUID)
			writer.PushVrI(int32(player.Gamemode))
			continue

		case UPDATE_LATENCY:
			writer.PushUID(player.Profile.UUID)
			writer.PushVrI(player.Ping)
			continue

		case UPDATE_DISPLAY_NAME:
			writer.PushUID(player.Profile.UUID)
			if player.Name == nil {
				writer.PushBit(false)
			} else {
				writer.PushBit(true)
				writer.PushTxt(player.Name.AsJson())
			}
			continue

		case REMOVE_PLAYER:
			writer.PushUID(player.Profile.UUID)
			continue
		}
	}
}
