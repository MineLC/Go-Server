package prot

import (
	"log"

	"github.com/golangmc/minecraft-server/apis/logs"
	"github.com/golangmc/minecraft-server/apis/task"
	"github.com/golangmc/minecraft-server/apis/util"
	"github.com/golangmc/minecraft-server/impl/base"
	"github.com/golangmc/minecraft-server/impl/game/mode"
	"github.com/golangmc/minecraft-server/impl/prot/server"
	client "github.com/golangmc/minecraft-server/impl/protocol/client"
)

type packets struct {
	util.Watcher

	logger *logs.Logging

	join chan base.PlayerAndConnection
	quit chan base.PlayerAndConnection
}

func NewPackets(tasking *task.Tasking, join chan base.PlayerAndConnection, quit chan base.PlayerAndConnection) base.Packets {
	packets := &packets{
		Watcher: util.NewWatcher(),

		logger: logs.NewLogging("protocol", logs.EveryLevel...),
	}

	mode.HandleState0(packets)
	mode.HandleState1(packets)
	mode.HandleState2(packets, join)
	mode.HandleState3(packets, packets.logger, tasking, join, quit)

	return packets
}

func (p *packets) GetPacketI(uuid int32, state base.PacketState) base.PacketI {
	creator := mapPackets[state][uuid]
	if creator == nil {
		log.Println("Found unknown packet: ")
		log.Println(uuid)
		log.Println(state)
		return nil
	}

	return creator()
}

var mapPackets map[base.PacketState]map[int32]func() base.PacketI = map[base.PacketState]map[int32]func() base.PacketI{
	base.SHAKE: {
		0: func() base.PacketI { return &server.PacketHandshakingInSetProtocol{} },
	},
	base.STATUS: {
		0: func() base.PacketI { return &server.PacketStatusInStart{} },
		1: func() base.PacketI { return &server.PacketStatusInPing{} },
	},
	base.LOGIN: {
		0: func() base.PacketI { return &server.PacketLoginInStart{} },
		1: func() base.PacketI { return &server.PacketLoginInEncryptionBegin{} },
		2: func() base.PacketI { return &server.PacketILoginPluginResponse{} },
	},
	base.PLAY: {
		0: func() base.PacketI { return &client.PacketPlayInKeepAlive{} },
		1: func() base.PacketI { return &client.PacketPlayInChatMessage{} },
		2: func() base.PacketI { return &client.PacketPlayInUseEntity{} },
		3: func() base.PacketI { return &client.PacketPlayInFlying{} },
		4: func() base.PacketI { return &client.PacketPlayInPosition{} },
		5: func() base.PacketI { return &client.PacketPlayInLook{} },
		6: func() base.PacketI { return &client.PacketPlayInPositionLook{} },
		// TODO: 7 - PacketPlayInBlockDig
		// TODO: 8 - PacketPlayInBlockPlace
		9:  func() base.PacketI { return &client.PacketPlayInHeldItemSlot{} },
		10: func() base.PacketI { return &client.PacketPlayInArmAnimation{} },
		11: func() base.PacketI { return &client.PacketPlayInEntityAction{} },
		// TODO: 12 - PacketPlayInSteerVehicle
		13: func() base.PacketI { return &client.PacketPlayInCloseWindow{} },
		14: func() base.PacketI { return &client.PacketPlayInWindowClick{} },
		15: func() base.PacketI { return &client.PacketPlayInTransaction{} },
		16: func() base.PacketI { return &client.PacketPlayInSetCreativeSlot{} },
		17: func() base.PacketI { return &client.PacketPlayInEnchantItem{} },
		// TODO: 18 - PacketPlayInUpdateSign
		19: func() base.PacketI { return &client.PacketPlayInAbilities{} },
		20: func() base.PacketI { return &client.PacketPlayInTabComplete{} },
		21: func() base.PacketI { return &client.PacketPlayInSettings{} },
		22: func() base.PacketI { return &client.PacketPlayInClientCommand{} },
		23: func() base.PacketI { return &client.PacketPlayInCustomPayload{} },
		24: func() base.PacketI { return &client.PacketPlayInSpectate{} },
		// TODO: 25 - PacketPlayInResourcePackStatus
	},
}
