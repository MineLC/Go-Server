package chunks

import (
	"math/bits"

	"github.com/minelc/go-server-api/network"
)

type PacketPlayOutChunkData struct {
	Load  bool
	Chunk Chunk
}

func (p *PacketPlayOutChunkData) UUID() int32 {
	return 33
}

func (p *PacketPlayOutChunkData) Push(writer network.Buffer) {
	writer.PushI32(p.Chunk.X)
	writer.PushI32(p.Chunk.Z)

	writer.PushBit(p.Load)
	chunkMap := createChunkMap(&p.Chunk)

	writer.PushI16(int16(chunkMap.size & '\uffff'))
	writer.PushUAS(chunkMap.bytes, true)
}

func getSize(i int) int {
	return i<<13 + i<<12 + 256
}

type ChunkMap struct {
	bytes []byte
	size  uint32
}

func createChunkMap(chunk *Chunk) *ChunkMap {
	chunkMap := ChunkMap{size: 0}
	validSections := [16]*ChunkSection{}
	sectionsAmount := 0

	for i := 0; i < 16; i++ {
		section := chunk.Sections[i]
		if section != nil {
			chunkMap.size |= 1 << i
			validSections[sectionsAmount] = section
			sectionsAmount++
		}
	}

	chunkMap.bytes = make([]byte, getSize(bits.OnesCount32(chunkMap.size)))
	bPos := 0

	for i := 0; i < 16; i++ {
		section := validSections[i]
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
		section := validSections[i]
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
