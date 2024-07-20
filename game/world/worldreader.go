package world

import (
	"bufio"
	"os"
	"strings"

	"github.com/DataDog/zstd"
	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/data"
	api_world "github.com/minelc/go-server-api/game/world"
	api_chunk "github.com/minelc/go-server-api/game/world/chunks"
	"github.com/minelc/go-server/game/world/chunks"
	"github.com/minelc/go-server/network"
	uuid "github.com/satori/go.uuid"
)

func readWorlds(manager *WorldManager) {
	files, err := os.ReadDir("worlds")
	if err != nil {
		api.GetServer().GetConsole().SendMsgColor("&cNo worlds found!")
		return
	}
	for _, file := range files {
		fileName := file.Name()
		worldName := getMinWorldName(fileName)
		if worldName == "" {
			continue
		}
		world := readWorld("worlds/"+fileName, worldName)
		if world == nil {
			continue
		}
		manager.worlds[worldName] = world
		api.GetServer().GetConsole().SendMsgColor("&bWorld loaded: " + worldName)
	}
}

func getMinWorldName(file string) string {
	if !strings.Contains(file, ".minworld") {
		api.GetServer().GetConsole().SendMsgColor("&cUnknown file in worlds folder: ", file)
		return ""
	}
	return file[0 : len(file)-9]
}

func readWorld(filePath string, worldName string) api_world.World {
	file, err := os.Open(filePath)
	if err != nil {
		api.GetServer().GetConsole().SendMsgColor("&cCan't open the world: &6"+worldName, err.Error())
		return nil
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		api.GetServer().GetConsole().SendMsgColor("&cError getting file info ", err.Error())
		return nil
	}
	reader := bufio.NewReader(file)
	version, chunkBufSize := getHeader(reader)

	if version != 1 {
		api.GetServer().GetConsole().SendMsgColor("&cOnly support v1")
		return nil
	}

	chunkBuf := make([]byte, chunkBufSize)
	chunkDataCompress := make([]byte, info.Size()-5) // 5 = Header size
	reader.Read(chunkDataCompress)

	zstd.Decompress(chunkBuf, chunkDataCompress)

	chunkDataCompress = nil
	buf := network.NewBufferWith(chunkBuf)
	amountChunks := buf.PullI32()
	world := World{
		Chunks: make(map[uint64]api_chunk.Chunk, amountChunks),
		Name:   worldName,
		Uuid:   uuid.UUID(data.CreateUUID(worldName)),
	}

	var i int32
	for i = 0; i < amountChunks; i++ {
		x := buf.PullI32()
		z := buf.PullI32()

		chunk := chunks.Chunk{X: x, Z: z}
		amountSections := int(buf.PullByt())
		for s := 0; s < amountSections; s++ {
			section := buf.PullByt()
			chunkSection := api_chunk.ChunkSection{}

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
