package chunks

import (
	block "github.com/minelc/go-server-api/data/block"
	api_chunk "github.com/minelc/go-server-api/game/world/chunks"
)

type Chunk struct {
	X int32
	Z int32

	Sections [16]*api_chunk.ChunkSection
}

func (c *Chunk) GetX() int32 {
	return c.X
}

func (c *Chunk) GetZ() int32 {
	return c.Z
}

func (c *Chunk) GetSections() [16]*api_chunk.ChunkSection {
	return c.Sections
}

func (c *Chunk) SetBlock(x int32, y int32, z int32, blockType block.Block, data uint16) {
	section := y >> 4
	cS := c.Sections[section]
	if cS == nil {
		cS = &api_chunk.ChunkSection{}
		c.Sections[section] = cS
	}

	x = x & 15
	z = z & 15
	y = y & 15

	key := (y << 8) | (z << 4) | x

	if blockType == 0 {
		if cS.Blocks[key] != 0 { // Change a normal block with air
			cS.BlocksPlaced--
		}
	} else {
		cS.BlocksPlaced++
	}
	cS.Blocks[key] = uint16(blockType)<<4 | data
}

func (c *Chunk) SetAll(blockType block.Block, start int, end uint16, data uint16) {
	block := uint16(blockType)<<4 | data
	for i := start; i < int(end); i++ {
		section := c.Sections[i]
		if section == nil {
			section = &api_chunk.ChunkSection{}
			c.Sections[i] = section
		}
		if block != 0 {
			section.BlocksPlaced += 4096
		}
		for b := 0; b < 4096; b++ {
			section.Blocks[b] = block
		}
	}
}
