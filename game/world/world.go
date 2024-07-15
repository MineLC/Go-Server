package world

import (
	"github.com/minelc/go-server/game/world/chunks"
)

type World struct {
	Chunks map[uint64]*chunks.Chunk
}

func compress(chunkX int32, chunkZ int32) uint64 {
	// Only for positive numbers:
	// First 32 bits = X
	// Second 32 bits = Z

	// To avoid collsion with negative numbers. Transform negative to positive
	var keyX, keyZ uint64
	if chunkX < 0 {
		keyX = 4294967199 - uint64(chunkX)
	} else {
		keyX = uint64(chunkX)
	}
	if chunkZ < 0 {
		keyZ = 4294967199 - uint64(chunkZ)
	} else {
		keyZ = uint64(chunkZ)
	}
	return (keyX << 32) | keyZ
}

/*
A simple guide to get chunk values:

chunkX := x >> 4
chunkZ := z >> 4
chunkSection := y >> 4
*/

func (w *World) GetChunk(chunkX int32, chunkZ int32) *chunks.Chunk {
	return w.Chunks[compress(chunkX, chunkZ)]
}

func (w *World) SetChunk(chunkX int32, chunkZ int32, chunk *chunks.Chunk) {
	if chunk == nil {
		return
	}
	w.Chunks[compress(chunkX, chunkZ)] = chunk
}

func (w *World) UnloadChunk(chunkX int32, chunkZ int32) {
	delete(w.Chunks, compress(chunkX, chunkZ))
}

func (w *World) GetAllChunks() []*chunks.Chunk {
	values := make([]*chunks.Chunk, 0, len(w.Chunks))
	for _, chunk := range w.Chunks {
		values = append(values, chunk)
	}
	return values
}
