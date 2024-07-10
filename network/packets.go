package network

import (
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/client/login"
	"github.com/minelc/go-server-api/network/client/play"
	"github.com/minelc/go-server-api/network/client/status"

	handler_login "github.com/minelc/go-server/network/client/login"
	handler_play "github.com/minelc/go-server/network/client/play"
	handler_status "github.com/minelc/go-server/network/client/status"
)

var playFuncs [25]func(c *network.Connection, packet network.PacketI)

func setDefaultHandlers() {
	playFuncs[0] = handler_play.HandleKeepAlive
	playFuncs[1] = handler_play.HandleChat
	playFuncs[20] = handler_play.HandleTab
}

func getPacketI(id int32, state network.PacketState) (network.PacketI, func(c *network.Connection, packet network.PacketI)) {
	switch state {
	case network.SHAKE:
		if id == 0 {
			return &status.PacketHandshakingInSetProtocol{}, handler_status.HandleHandShake
		}
		return nil, nil

	case network.STATUS:
		if id == 0 {
			return &status.PacketIRequest{}, nil
		}
		if id == 1 {
			return &status.PacketIPing{}, handler_status.HandlePing
		}
		return nil, nil

	case network.LOGIN:
		switch id {
		case 0:
			return &login.PacketILoginStart{}, handler_login.HandleLoginStart
		case 1:
			return &login.PacketIEncryptionResponse{}, handler_login.HandleEncryption
		case 2:
			return &login.PacketILoginPluginResponse{}, nil
		}
		return nil, nil

	case network.PLAY:
		switch id {
		case 0:
			return &play.PacketPlayInKeepAlive{}, playFuncs[0]
		case 1:
			return &play.PacketPlayInChatMessage{}, playFuncs[1]
		case 2:
			return &play.PacketPlayInUseEntity{}, playFuncs[2]
		case 3:
			return &play.PacketPlayInFlying{}, playFuncs[3]
		case 4:
			return &play.PacketPlayInPosition{}, playFuncs[4]
		case 5:
			return &play.PacketPlayInLook{}, playFuncs[5]
		case 6:
			return &play.PacketPlayInPositionLook{}, playFuncs[6]
		case 7:
			return &play.PacketPlayInBlockDig{}, playFuncs[7]
		// 8 TODO: PacketPlayInBlockPlace
		case 9:
			return &play.PacketPlayInHeldItemSlot{}, playFuncs[9]
		case 10:
			return &play.PacketPlayInArmAnimation{}, playFuncs[10]
		case 11:
			return &play.PacketPlayInEntityAction{}, playFuncs[11]
		// 12 TODO: PacketPlayInSteerVehicle
		case 13:
			return &play.PacketPlayInCloseWindow{}, playFuncs[13]
		case 14:
			return &play.PacketPlayInWindowClick{}, playFuncs[14]
		case 15:
			return &play.PacketPlayInTransaction{}, playFuncs[15]
		case 16:
			return &play.PacketPlayInSetCreativeSlot{}, playFuncs[16]
		case 17:
			return &play.PacketPlayInEnchantItem{}, playFuncs[17]
		// 18 TODO: PacketPlayInUpdateSign]
		case 19:
			return &play.PacketPlayInAbilities{}, playFuncs[19]
		case 20:
			return &play.PacketPlayInTabComplete{}, playFuncs[20]
		case 21:
			return &play.PacketPlayInSettings{}, playFuncs[21]
		case 22:
			return &play.PacketPlayInClientCommand{}, playFuncs[22]
		case 23:
			return &play.PacketPlayInCustomPayload{}, playFuncs[23]
		case 24:
			return &play.PacketPlayInSpectate{}, playFuncs[24]
		}
		return nil, nil
	}

	return nil, nil
}
