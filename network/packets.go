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

type Packets struct {
	playFuncs   [25]func(c network.Connection, packet network.PacketI)
	compression int
}

func NewDefaultHandler(compressionThreshold int) network.PacketManager {
	p := Packets{compression: compressionThreshold}
	p.playFuncs[network.KeepAlive] = handler_play.HandleKeepAlive
	p.playFuncs[network.ChatMessage] = handler_play.HandleChat
	p.playFuncs[network.TabComplete] = handler_play.HandleTab
	p.playFuncs[network.Position] = handler_play.HandleMove
	p.playFuncs[network.Look] = handler_play.HandleLook
	p.playFuncs[network.PositionLook] = handler_play.HandleMoveLook
	p.playFuncs[network.Settings] = handler_play.HandleSettings
	return &p
}

func (p *Packets) RemoveHandler(id network.PacketInput) {
	if id < 0 || id > 24 {
		return
	}
	p.playFuncs[id] = nil
}

func (p *Packets) GetCompression() int {
	return p.compression
}

func (p *Packets) SetHandler(id network.PacketInput, handler func(c network.Connection, packet network.PacketI)) {
	p.playFuncs[id] = handler
}

func (p *Packets) GetPlayFunc(id int32) func(c network.Connection, packet network.PacketI) {
	if id < 0 || id > 24 {
		return nil
	}
	return p.playFuncs[id]
}

func (p *Packets) getPacketI(id int32, state network.PacketState) (network.PacketI, func(c network.Connection, packet network.PacketI)) {
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
			return &play.PacketPlayInKeepAlive{}, p.playFuncs[0]
		case 1:
			return &play.PacketPlayInChatMessage{}, p.playFuncs[1]
		case 2:
			return &play.PacketPlayInUseEntity{}, p.playFuncs[2]
		case 3:
			return &play.PacketPlayInFlying{}, p.playFuncs[3]
		case 4:
			return &play.PacketPlayInPosition{}, p.playFuncs[4]
		case 5:
			return &play.PacketPlayInLook{}, p.playFuncs[5]
		case 6:
			return &play.PacketPlayInPositionLook{}, p.playFuncs[6]
		case 7:
			return &play.PacketPlayInBlockDig{}, p.playFuncs[7]
		// 8 TODO: PacketPlayInBlockPlace
		case 9:
			return &play.PacketPlayInHeldItemSlot{}, p.playFuncs[9]
		case 10:
			return &play.PacketPlayInArmAnimation{}, p.playFuncs[10]
		case 11:
			return &play.PacketPlayInEntityAction{}, p.playFuncs[11]
		// 12 TODO: PacketPlayInSteerVehicle
		case 13:
			return &play.PacketPlayInCloseWindow{}, p.playFuncs[13]
		case 14:
			return &play.PacketPlayInWindowClick{}, p.playFuncs[14]
		case 15:
			return &play.PacketPlayInTransaction{}, p.playFuncs[15]
		case 16:
			return &play.PacketPlayInSetCreativeSlot{}, p.playFuncs[16]
		case 17:
			return &play.PacketPlayInEnchantItem{}, p.playFuncs[17]
		// 18 TODO: PacketPlayInUpdateSign]
		case 19:
			return &play.PacketPlayInAbilities{}, p.playFuncs[19]
		case 20:
			return &play.PacketPlayInTabComplete{}, p.playFuncs[20]
		case 21:
			return &play.PacketPlayInSettings{}, p.playFuncs[21]
		case 22:
			return &play.PacketPlayInClientCommand{}, p.playFuncs[22]
		case 23:
			return &play.PacketPlayInCustomPayload{}, p.playFuncs[23]
		case 24:
			return &play.PacketPlayInSpectate{}, p.playFuncs[24]
		}
		return nil, nil
	}

	return nil, nil
}
