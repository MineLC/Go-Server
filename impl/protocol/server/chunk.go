package server

import (
	"github.com/golangmc/minecraft-server/apis/buff"
	"github.com/golangmc/minecraft-server/apis/game/level"
	"github.com/golangmc/minecraft-server/impl/base"
	apis_conn "github.com/golangmc/minecraft-server/impl/conn"
)

type PacketPlayOutMapChunk struct {
	Chunk level.Chunk
}

func (p *PacketPlayOutMapChunk) UUID() int32 {
	return 33
}

func (p *PacketPlayOutMapChunk) Push(writer buff.Buffer, conn base.Connection) {
	writer.PushI32(int32(p.Chunk.ChunkX()))
	writer.PushI32(int32(p.Chunk.ChunkZ()))

	// full chunk (for now >:D)
	writer.PushBit(true)

	chunkData := apis_conn.NewBuffer()
	p.Chunk.Push(chunkData) // write chunk data and primary bit mask

	// pull primary bit mask and push to writer
	writer.PushVrI(chunkData.PullVrI())

	// write height-maps
	writer.PushNbt(p.Chunk.HeightMapNbtCompound())

	biomes := make([]int32, 1024, 1024)
	for i := range biomes {
		biomes[i] = 0 // void biome
	}

	for _, biome := range biomes {
		writer.PushI32(biome)
	}

	// data, prefixed with len
	writer.PushUAS(chunkData.UAS(), true)

	// write block entities
	writer.PushVrI(0)
}
