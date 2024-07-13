package generator

import (
	block "github.com/minelc/go-server-api/data/block"
	"github.com/minelc/go-server/game/world/chunks"
)

func GenerateFlatChunk(x int32, z int32) *chunks.Chunk {
	chunk := chunks.Chunk{X: x, Z: z}
	section := &chunks.ChunkSection{}

	blockType := block.BEDROCK

	for y := 0; y < 4; y++ {
		blockData := uint16(blockType << 4)
		yCord := y << 8

		for x := 0; x < 16; x++ {
			for z := 0; z < 16; z++ {
				section.Blocks[yCord|(z<<4)|x] = blockData
			}
		}
		if y == 0 {
			blockType = block.DIRT
			continue
		}
		if y == 2 {
			blockType = block.GRASS
		}
	}
	chunk.Sections[0] = section
	return &chunk
}
