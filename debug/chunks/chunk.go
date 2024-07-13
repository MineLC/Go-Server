package chunks

import block "github.com/minelc/go-server-api/data/block"

type ChunkSection struct {
	Blocks [4096]uint16
}

type Chunk struct {
	X int32
	Y int32

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

	cS.Blocks[(y<<8)|(z<<4)|x] = uint16(blockType)<<4 | data
}

func (c *Chunk) SetAll(blockType block.Block, section int, data uint16) {
	block := uint16(blockType)<<4 | data
	for i := 0; i < section; i++ {
		section := c.Sections[i]
		if section == nil {
			section = &ChunkSection{}
			c.Sections[i] = section
		}

		for b := 0; b < 4096; b++ {
			section.Blocks[b] = block
		}
	}
}
