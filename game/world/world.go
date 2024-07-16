package world

import (
	chunks "github.com/minelc/go-server-api/game/world/chunks"
)

type World struct {
	Chunks map[uint64]chunks.Chunk
}

/*
A simple guide to get chunk values:

chunkX := x >> 4
chunkZ := z >> 4
chunkSection := y >> 4
*/

func (w *World) GetChunk(chunkX int32, chunkZ int32) chunks.Chunk {
	return w.Chunks[compress(chunkX, chunkZ)]
}

func (w *World) SetChunk(chunkX int32, chunkZ int32, chunk chunks.Chunk) {
	if chunk == nil {
		return
	}
	w.Chunks[compress(chunkX, chunkZ)] = chunk
}

func (w *World) UnloadChunk(chunkX int32, chunkZ int32) {
	delete(w.Chunks, compress(chunkX, chunkZ))
}

func (w *World) GetAllChunks() []chunks.Chunk {
	values := make([]chunks.Chunk, 0, len(w.Chunks))
	for _, chunk := range w.Chunks {
		values = append(values, chunk)
	}
	return values
}
