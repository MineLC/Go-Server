package mode

import (
	"time"

	"github.com/golangmc/minecraft-server/apis"
	"github.com/golangmc/minecraft-server/apis/data"
	"github.com/golangmc/minecraft-server/apis/data/chat"
	"github.com/golangmc/minecraft-server/apis/data/msgs"
	"github.com/golangmc/minecraft-server/apis/game"
	"github.com/golangmc/minecraft-server/apis/logs"
	"github.com/golangmc/minecraft-server/apis/task"
	"github.com/golangmc/minecraft-server/apis/util"
	"github.com/golangmc/minecraft-server/impl/base"
	apis_conn "github.com/golangmc/minecraft-server/impl/conn"
	"github.com/golangmc/minecraft-server/impl/data/client"
	"github.com/golangmc/minecraft-server/impl/data/plugin"

	impl_event "github.com/golangmc/minecraft-server/impl/game/event"

	server_packet "github.com/golangmc/minecraft-server/impl/prot/server"
	client_packets "github.com/golangmc/minecraft-server/impl/protocol/client"
	server_packets "github.com/golangmc/minecraft-server/impl/protocol/server"
)

func HandleState3(watcher util.Watcher, logger *logs.Logging, tasking *task.Tasking, join chan base.PlayerAndConnection, quit chan base.PlayerAndConnection) {

	tasking.EveryTime(10, time.Second, func(task *task.Task) {

		api := apis.MinecraftServer()

		// I hate this, add a functional method for player iterating
		for _, player := range api.Players() {

			// also probably add one that returns both the player and their connection
			conn := api.ConnByUUID(player.UUID())
			// keep player connection alive via keep alive
			conn.SendPacket(&server_packets.PacketPlayOutKeepAlive{KeepAliveID: int32(time.Now().UnixNano() / 1000000)})
		}
		/*
					level := impl_level.NewLevel("test")
			impl_level.GenSuperFlat(level, 6)

			for _, chunk := range level.Chunks() {
				conn.SendPacket(&client_packet.PacketOChunkData{Chunk: chunk})
			}

			logger.DataF("chunks sent to player: %s", conn.Player.Name())

		*/
	})

	watcher.SubAs(func(packet *client_packets.PacketPlayInKeepAlive, conn base.Connection) {
		logger.DataF("player %s is being kept alive", conn.Address())
	})

	watcher.SubAs(func(packet *server_packet.PacketIPluginMessage, conn base.Connection) {
		api := apis.MinecraftServer()

		player := api.PlayerByConn(conn)
		if player == nil {
			return // log no player found?
		}

		api.Watcher().PubAs(impl_event.PlayerPluginMessagePullEvent{
			Conn: base.PlayerAndConnection{
				Connection: conn,
				Player:     player,
			},
			Channel: packet.Message.Chan(),
			Message: packet.Message,
		})
	})

	watcher.SubAs(func(packet *client_packets.PacketPlayInChatMessage, conn base.Connection) {
		api := apis.MinecraftServer()

		who := api.PlayerByConn(conn)

		if packet.Message[0] == '/' {
			who.SendMessage("Ejecutaste un comando xd. ")
			who.SendMessage(packet.Message[1:len(packet.Message)])
			return
		}

		out := msgs.
			New(who.Name()).SetColor(chat.White).
			Add(":").SetColor(chat.Gray).
			Add(" ").
			Add(chat.Translate(packet.Message)).SetColor(chat.White).
			AsText() // why not just use translate?

		api.Broadcast(out)
	})

	go func() {
		for conn := range join {
			apis.MinecraftServer().Watcher().PubAs(impl_event.PlayerConnJoinEvent{Conn: conn})

			conn.SendPacket(&server_packets.PacketPlayOutLogin{
				EntityID:    int32(conn.EntityUUID()),
				Hardcore:    false,
				GameMode:    game.CREATIVE,
				Dimension:   game.OVERWORLD,
				Difficulty:  game.EASY,
				MaxPlayers:  10,
				LevelType:   game.DEFAULT,
				ReduceDebug: false,
			})

			conn.SendPacket(&server_packets.PacketPlayOutCustomPayload{
				Message: &plugin.Brand{
					Name: chat.Translate("MC|Brand"),
				},
				Buffer: apis_conn.NewBufferWith([]byte("LC")),
			})
			conn.SendPacket(&server_packets.PacketPlayOutServerDifficulty{
				Difficulty: game.PEACEFUL,
			})
			red := chat.DarkRed
			blue := chat.Blue
			conn.SendPacket(&server_packets.PacketPlayOutTabInfo{
				Header: msgs.Message{
					Text:  "este es mi header",
					Color: &red,
				},
				Footer: msgs.Message{
					Text:  "este es mi footer",
					Color: &blue,
				},
			})
			conn.SendPacket(&server_packets.PacketPlayOutAbilities{
				Abilities: client.PlayerAbilities{
					Invulnerable: true,
					Flying:       true,
					AllowFlight:  true,
					InstantBuild: false,
				},
				FlyingSpeed: 0.05, // default value
				FieldOfView: 0.1,  // default value
			})

			conn.SendPacket(&server_packets.PacketPlayOutHeldItemChange{
				Slot: client.SLOT_0,
			})

			conn.SendPacket(&server_packets.PacketPlayOutPosition{
				Location: data.Location{
					PositionF: data.PositionF{
						X: 0,
						Y: 10,
						Z: 0,
					},
					RotationF: data.RotationF{
						AxisX: 0,
						AxisY: 0,
					},
				},
				Relative: client.Relativity{},
			})

			conn.SendPacket(&server_packets.PacketPlayOutEntityMetadata{Entity: conn.Player})

			conn.SendPacket(&server_packets.PacketPlayOutPosition{
				Location: data.Location{
					PositionF: data.PositionF{
						X: 0,
						Y: 10,
						Z: 0,
					},
					RotationF: data.RotationF{
						AxisX: 0,
						AxisY: 0,
					},
				},
				Relative: client.Relativity{},
			})
		}
	}()

	go func() {
		for conn := range quit {
			apis.MinecraftServer().Watcher().PubAs(impl_event.PlayerConnQuitEvent{Conn: conn})
		}
	}()
}
