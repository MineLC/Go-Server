package chunks

import (
	"math/bits"

	"github.com/minelc/go-server-api/network"
)

type PacketPlayOutChunkData struct {
	// To unload a chunk set GroundUp in true
	GroundUp bool

	Chunk Chunk
}

func (p *PacketPlayOutChunkData) UUID() int32 {
	return 33
}

func (p *PacketPlayOutChunkData) Push(writer network.Buffer) {
	writer.PushI32(p.Chunk.X)
	writer.PushI32(p.Chunk.Y)
	flag := true

	writer.PushBit(flag)
	chunkMap := ab(&p.Chunk, flag, 0)

	writer.PushI16(int16(chunkMap.size & '\uffff'))
	writer.PushUAS(chunkMap.bytes, true)
}

func a(i int) int {
	j := i * 2 * 16 * 16 * 16
	k := i * 16 * 16 * 16 / 2
	l := i * 16 * 16 * 16 / 2
	return j + k + l + 256
}

type ChunkMap struct {
	bytes []byte
	size  uint32
}

func ab(chunk *Chunk, flag bool, i int) *ChunkMap {
	sections := chunk.Sections
	chunkMap := ChunkMap{size: 0}
	validSections := [16]*ChunkSection{}
	position := 0

	for j := 0; j < 16; j++ {
		section := sections[i]
		chunkMap.size |= 1 << j
		if section != nil {
			validSections[position] = section
			position++
		}
	}

	chunkMap.bytes = make([]byte, a(bits.OnesCount32(chunkMap.size)))
	bPos := 0

	for i := 0; i < 16; i++ {
		section := validSections[i]
		if section == nil {
			continue
		}
		for j := 0; j < 4096; j++ {
			c0 := section.Blocks[j]
			chunkMap.bytes[bPos] = (byte)(c0 & 255)
			bPos++
			chunkMap.bytes[bPos] = (byte)(c0 >> 8 & 255)
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

	/*
		Set Sky light
		if flag1 {
			for i := 0; i < 16; i++ {
				for j := 0; j < 2048; j++ {
					chunkMap.bytes[bPos] = chunk.SkyLight[j]
					bPos++
				}
			}
		}
	*/
	/*
		Set biome
		if flag {
			for i := 0; i < 256; i++ {
				chunkMap.bytes[bPos] = chunk.Biomes[i]
				bPos++
			}
		}
	*/

	return &chunkMap
}
