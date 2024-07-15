package chunks

import (
	"math/bits"

	"github.com/minelc/go-server-api/network"
)

type PacketPlayOutChunkData struct {
	Chunk *Chunk
}

func (p *PacketPlayOutChunkData) UUID() int32 {
	return 33
}

func (p *PacketPlayOutChunkData) Push(writer network.Buffer) {
	writer.PushI32(p.Chunk.X)
	writer.PushI32(p.Chunk.Z)

	writer.PushBit(true)

	chunkMap := createChunkMap(p.Chunk, true, true, '\uffff')
	writer.PushI16(int16(chunkMap.size & '\uffff'))
	writer.PushUAS(chunkMap.bytes, true)
}

type PacketPlayOutUnloadChunk struct {
	X int32
	Z int32
}

func (p *PacketPlayOutUnloadChunk) UUID() int32 {
	return 33
}

func (p *PacketPlayOutUnloadChunk) Push(writer network.Buffer) {
	writer.PushI32(p.X)
	writer.PushI32(p.Z)

	writer.PushBit(true)

	writer.PushI16(0)
	writer.PushVrI(0)
}

type ChunkMap struct {
	bytes []byte
	size  uint32
}

func size(i int, flag bool, flag1 bool) int {
	j := i * 2 * 16 * 16 * 16
	k := i * 16 * 16 * 16 / 2
	var l, i1 int
	if flag {
		l = i * 16 * 16 * 16 / 2
	}
	if flag1 {
		i1 = 256
	}
	return j + k + l + i1
}

func createChunkMap(chunk *Chunk, flag bool, skyLight bool, mask uint16) *ChunkMap {
	chunkMap := ChunkMap{size: 0}

	for i := 0; i < 16; i++ {
		section := chunk.Sections[i]
		if section != nil && section.BlocksPlaced != 0 && mask&(1<<i) != 0 {
			chunkMap.size |= 1 << i
		}
	}

	chunkMap.bytes = make([]byte, size(bits.OnesCount32(chunkMap.size), skyLight, flag))
	bPos := 0

	for i := 0; i < 16; i++ {
		section := chunk.Sections[i]
		if section == nil {
			continue
		}
		for j := 0; j < 4096; j++ {
			c0 := section.Blocks[j]
			chunkMap.bytes[bPos] = (byte)(c0 & 255) // Block id
			bPos++
			chunkMap.bytes[bPos] = (byte)(c0 >> 8 & 255) // Block data
			bPos++
		}
	}

	// Set light for all blocks
	for i := 0; i < 16; i++ {
		section := chunk.Sections[i]
		if section == nil {
			continue
		}
		for j := 0; j < 2048; j++ {
			chunkMap.bytes[bPos] = 255
			bPos++
		}
	}
	return &chunkMap
}
