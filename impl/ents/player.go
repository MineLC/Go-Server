package ents

import (
	"github.com/minelc/go-server/api"
	"github.com/minelc/go-server/api/data"
	"github.com/minelc/go-server/api/data/chat"
	player_data "github.com/minelc/go-server/api/data/player"
	"github.com/minelc/go-server/api/ents"
	"github.com/minelc/go-server/api/network"
	play "github.com/minelc/go-server/impl/network/server/play"
)

type player struct {
	entityLiving

	prof     *player_data.Profile
	online   bool
	uuid     data.UUID
	gamemode player_data.GameMode
	conn     network.Connection
	ping     int32

	keep_alive_delay int64
}

func NewPlayer(prof *player_data.Profile, conn network.Connection) ents.Player {
	player := &player{
		prof:             prof,
		entityLiving:     NewEntityLiving(),
		gamemode:         player_data.CREATIVE,
		keep_alive_delay: -1,
	}

	player.name = prof.Name
	player.uuid = prof.UUID

	player.SetConn(conn)

	return player
}

func (p *player) SendMsg(messages ...string) {
	p.SendMsgPos(chat.NormalChat, messages...)
}

func (p *player) SendMsgColor(messages ...string) {
	p.SendMsgColorPos(chat.NormalChat, messages...)
}

func (p *player) SendMsgColorPos(pos chat.MessagePosition, messages ...string) {
	for _, msg := range messages {
		p.conn.SendPacket(&play.PacketPlayOutChatMessage{
			Message:         *chat.New(chat.Translate(msg)),
			MessagePosition: pos,
		})
	}
}

func (p *player) SendMsgPos(pos chat.MessagePosition, messages ...string) {
	for _, msg := range messages {
		p.conn.SendPacket(&play.PacketPlayOutChatMessage{
			Message:         *chat.New(msg),
			MessagePosition: pos,
		})
	}
}

func (p *player) GetIsOnline() bool {
	return p.online
}

func (p *player) SetIsOnline(state bool) {
	p.online = state
}

func (p *player) GetProfile() *player_data.Profile {
	return p.prof
}

func (p *player) SetConn(conn network.Connection) {
	p.conn = conn
}

func (p *player) UUID() data.UUID {
	return p.uuid
}

func (p *player) GetGamemode() player_data.GameMode {
	return p.gamemode
}

func (p *player) SetGamemode(gamemode player_data.GameMode) {
	p.gamemode = gamemode
}

func (p *player) Disconnect() {
	api.GetServer().Disconnect(p.conn)

	p.conn.Stop()
}

func (p *player) GetPing() int32 {
	return p.ping
}

func (p *player) SetPing(ping_delay int64, server_ping int64) {
	p.ping = int32(ping_delay - server_ping)
}

func (p *player) GetKeepAlive() int64 {
	return p.keep_alive_delay
}

func (p *player) SetKeepAlive(time int64) {
	p.keep_alive_delay = time
}
