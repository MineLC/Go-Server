package ents

import (
	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/data/chat"
	player_data "github.com/minelc/go-server-api/data/player"
	"github.com/minelc/go-server-api/ents"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server-api/network/server/play"
	"github.com/minelc/go-server-api/plugin"
	"github.com/minelc/go-server-api/plugin/events"
	"github.com/minelc/go-server/game/join"
)

type player struct {
	entityLiving

	prof     *player_data.Profile
	online   bool
	gamemode player_data.GameMode
	conn     network.Connection
	ping     int32

	Settings ents.ClientSettings

	exp        int32
	absorption byte
	food       float32

	keep_alive_delay int64
}

func NewPlayer(prof *player_data.Profile, conn network.Connection) ents.Player {
	player := &player{
		prof:             prof,
		entityLiving:     NewEntityLiving(),
		gamemode:         player_data.CREATIVE,
		keep_alive_delay: -1,
		conn:             conn,
	}

	player.nametag = nil
	player.name = prof.Name

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

func (p *player) GetConnection() network.Connection {
	return p.conn
}

func (p *player) SetIsOnline(state bool) {
	p.online = state
}

func (p *player) GetProfile() *player_data.Profile {
	return p.prof
}

func (p *player) UUID() data.UUID {
	return p.prof.UUID
}

func (p *player) GetGamemode() player_data.GameMode {
	return p.gamemode
}

func (p *player) SetGamemode(gamemode player_data.GameMode) {
	p.gamemode = gamemode
}

func (p *player) Disconnect() {
	api.GetServer().Disconnect(p.conn)

	api.GetServer().GetPluginManager().CallEvent(events.PlayerQuitEvent{Player: p}, plugin.Quit)

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

func (p *player) GetAbsorption() byte {
	return p.absorption
}

func (p *player) GetIsOnline() bool {
	return p.online
}

func (p *player) GetClientSettings() *ents.ClientSettings {
	return &p.Settings
}

func (p *player) GetFood() float32 {
	return p.food
}

func (p *player) GetLevel() int32 {
	return join.GetLevel(int(p.exp))
}

func (p *player) GetXP() int32 {
	return p.exp
}

func (p *player) SetLevel(level int32) {
	p.exp = join.GetXP(level)
	p.conn.SendPacket(&join.PacketPlayOutExperience{Xp: p.exp, Level: level})
}

func (p *player) SetXP(xp int32) {
	p.exp = xp
	p.conn.SendPacket(&join.PacketPlayOutExperience{Xp: p.exp, Level: p.GetLevel()})
}

func (e *player) PushMetadata(buffer network.Buffer) {
	buffer.PushByt(HumanSkin)
	buffer.PushByt(e.Settings.SkinParts)
}
