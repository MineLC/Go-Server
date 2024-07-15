package file

import (
	"bufio"
	"os"

	"github.com/DataDog/zstd"
	"github.com/minelc/go-server/game/world"
	"github.com/minelc/go-server/game/world/chunks"
	"github.com/minelc/go-server/network"
)

func ReadWorld() *world.World {
	file, err := os.Open("/home/choco/Escritorio/server/1.8.8/plugins/MinimalWorld/worlds/testa.minworld")
	if err != nil {
		println("&cCan't open the world")
		//api.GetServer().GetConsole().SendMsgColor("&cCan't open the world: &e"+"testa.minworld", err.Error())
		return nil
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		println("&cError getting file info, " + err.Error())
		return nil
	}
	reader := bufio.NewReader(file)
	version, chunkBufSize := getHeader(reader)

	if version != 1 {
		println("&cOnly support v1")
		return nil
	}

	chunkBuf := make([]byte, chunkBufSize)
	chunkDataCompress := make([]byte, info.Size()-5) // 5 = Header size
	reader.Read(chunkDataCompress)

	zstd.Decompress(chunkBuf, chunkDataCompress)

	chunkDataCompress = nil
	buf := network.NewBufferWith(chunkBuf)
	amountChunks := buf.PullI32()
	world := world.World{Chunks: make(map[uint64]*chunks.Chunk, amountChunks)}

	var i int32
	for i = 0; i < amountChunks; i++ {
		x := buf.PullI32()
		z := buf.PullI32()

		chunk := chunks.Chunk{X: x, Z: z}
		amountSections := int(buf.PullByt())
		for s := 0; s < amountSections; s++ {
			section := buf.PullByt()
			chunkSection := chunks.ChunkSection{}

			for b := 0; b < 4096; b++ {
				block := buf.PullU16()
				chunkSection.Blocks[b] = block
				if block != 0 {
					chunkSection.BlocksPlaced++
				}
			}
			chunk.Sections[section] = &chunkSection
		}
		world.SetChunk(x, z, &chunk)
	}
	return &world
}

func getHeader(reader *bufio.Reader) (byte, int32) {
	header := make([]byte, 5)
	reader.Read(header)
	headerBuf := network.NewBufferWith(header)

	return headerBuf.PullByt(), headerBuf.PullI32()
}
