package network

import (
	"github.com/minelc/go-server/api/network"
	"github.com/minelc/go-server/impl/network/client/login"
	"github.com/minelc/go-server/impl/network/client/play"
	"github.com/minelc/go-server/impl/network/client/status"
)

func getPacketI(id int32, state network.PacketState) network.PacketI {
	switch state {
	case network.SHAKE:
		if id == 0 {
			return &status.PacketHandshakingInSetProtocol{}
		}
		return nil

	case network.STATUS:
		if id == 0 {
			return &status.PacketIRequest{}
		}
		if id == 1 {
			return &status.PacketIPing{}
		}
		return nil

	case network.LOGIN:
		switch id {
		case 0:
			return &login.PacketILoginStart{}
		case 1:
			return &login.PacketIEncryptionResponse{}
		case 2:
			return &login.PacketILoginPluginResponse{}
		}
		return nil

	case network.PLAY:
		switch id {
		case 0:
			return &play.PacketPlayInKeepAlive{}
		case 1:
			return &play.PacketPlayInChatMessage{}
		case 2:
			return &play.PacketPlayInUseEntity{}
		case 3:
			return &play.PacketPlayInFlying{}
		case 4:
			return &play.PacketPlayInPosition{}
		case 5:
			return &play.PacketPlayInLook{}
		case 6:
			return &play.PacketPlayInPositionLook{}
		// TODO: 7 - PacketPlayInBlockDig
		// TODO: 8 - PacketPlayInBlockPlace
		case 9:
			return &play.PacketPlayInHeldItemSlot{}
		case 10:
			return &play.PacketPlayInArmAnimation{}
		case 11:
			return &play.PacketPlayInEntityAction{}
		// TODO: 12 - PacketPlayInSteerVehicle
		case 13:
			return &play.PacketPlayInCloseWindow{}
		case 14:
			return &play.PacketPlayInWindowClick{}
		case 15:
			return &play.PacketPlayInTransaction{}
		case 16:
			return &play.PacketPlayInSetCreativeSlot{}
		case 17:
			return &play.PacketPlayInEnchantItem{}
		// TODO: 18 - PacketPlayInUpdateSign
		case 19:
			return &play.PacketPlayInAbilities{}
		case 20:
			return &play.PacketPlayInTabComplete{}
		case 21:
			return &play.PacketPlayInSettings{}
		case 22:
			return &play.PacketPlayInClientCommand{}
		case 23:
			return &play.PacketPlayInCustomPayload{}
		case 24:
			return &play.PacketPlayInSpectate{}
		}
		return nil
	}

	return nil
}
