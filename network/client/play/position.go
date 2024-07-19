package play

import (
	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/client/play"
)

func HandleMove(c network.Connection, packet network.PacketI) {
	p := packet.(*play.PacketPlayInPosition)
	pos := api.GetServer().GetPlayer(c).GetPosition()
	pos.X = p.X
	pos.Y = p.Y
	pos.Z = p.Z
}

func HandleLook(c network.Connection, packet network.PacketI) {
	p := packet.(*play.PacketPlayInLook)
	head := api.GetServer().GetPlayer(c).GetHeadPos()
	head.AxisX = p.Yaw
	head.AxisY = p.Pitch
}

func HandleMoveLook(c network.Connection, packet network.PacketI) {
	p := packet.(*play.PacketPlayInPositionLook)

	player := api.GetServer().GetPlayer(c)
	pos := player.GetPosition()
	pos.X = p.X
	pos.Y = p.Y
	pos.Z = p.Z

	head := player.GetHeadPos()
	head.AxisX = p.Yaw
	head.AxisY = p.Pitch
}
