package server

import "github.com/minelc/go-server/api/network"

type PacketPlayOutKeepAlive struct {
	KeepAliveID int32
}

func (p *PacketPlayOutKeepAlive) UUID() int32 {
	return 0
}

func (p *PacketPlayOutKeepAlive) Push(writer network.Buffer) {
	writer.PushVrI(p.KeepAliveID)
}
