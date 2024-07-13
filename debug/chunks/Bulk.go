package chunks

import "github.com/minelc/go-server-api/network"

type PacketPlayOutMapChunkBulk struct {
	// To unload a chunk set GroundUp in true

	Chunk Chunk
}

func (p *PacketPlayOutMapChunkBulk) UUID() int32 {
	return 38
}

func (p *PacketPlayOutMapChunkBulk) Push(writer network.Buffer) {
	writer.PushBit(false)
	writer.PushVrI(1)
	writer.PushI32(0)
	writer.PushI32(0)
	chunkMap := ab(&p.Chunk, true, '\uffff')
	writer.PushI16(int16(chunkMap.size & '\uffff'))
	writer.PushUAS(chunkMap.bytes, false)
}
