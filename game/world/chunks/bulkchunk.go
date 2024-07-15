package chunks

import "github.com/minelc/go-server-api/network"

type PacketPlayOutBulkChunkData struct {
	Chunks []*Chunk
}

func (p *PacketPlayOutBulkChunkData) UUID() int32 {
	return 38
}

func (p *PacketPlayOutBulkChunkData) Push(writer network.Buffer) {
	writer.PushBit(true)

	length := len(p.Chunks)
	writer.PushVrI(int32(length))
	chunksMaps := make([]*ChunkMap, length)

	for i := 0; i < length; i++ {
		chunk := p.Chunks[i]
		writer.PushI32(chunk.X)
		writer.PushI32(chunk.Z)

		chunkMap := createChunkMap(p.Chunks[i], true, true, '\uffff')
		writer.PushI16(int16(chunkMap.size & '\uffff'))
		chunksMaps[i] = chunkMap
	}

	for i := 0; i < length; i++ {
		writer.PushUAS(chunksMaps[i].bytes, false)
	}
}
