package chunks

import block "github.com/minelc/go-server-api/data/block"

type ChunkSection struct {
	Blocks       [4096]uint16
	BlocksPlaced int16
}

func (c *ChunkSection) GetBlocksPlaced() int16 {
	return c.BlocksPlaced
}

type Chunk struct {
	X int32
	Z int32

	Sections [16]*ChunkSection
}

func (c *Chunk) SetBlock(x int32, y int32, z int32, blockType block.Block, data uint16) {
	section := y >> 4
	cS := c.Sections[section]
	if cS == nil {
		cS = &ChunkSection{}
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

func (c *Chunk) SetAll(blockType block.Block, section int, data uint16) {
	block := uint16(blockType)<<4 | data
	for i := 0; i < section; i++ {
		section := c.Sections[i]
		if section == nil {
			section = &ChunkSection{}
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
