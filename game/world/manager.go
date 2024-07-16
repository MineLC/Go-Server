package world

import (
	"github.com/minelc/go-server-api/game/world"
)

type WorldManager struct {
	worlds       map[string]world.World
	defaultWorld world.World
}

func compress(chunkX int32, chunkZ int32) uint64 {
	// Only for positive numbers:
	// First 32 bits = X
	// Second 32 bits = Z

	// To avoid collsion with negative numbers. Transform negative to positive
	var keyX, keyZ uint64
	if chunkX < 0 {
		keyX = 4294967199 - uint64(chunkX)
	} else {
		keyX = uint64(chunkX)
	}
	if chunkZ < 0 {
		keyZ = 4294967199 - uint64(chunkZ)
	} else {
		keyZ = uint64(chunkZ)
	}
	return (keyX << 32) | keyZ
}

func (w *WorldManager) Compress(chunkX int32, chunkZ int32) uint64 {
	return compress(chunkX, chunkZ)
}

func (w *WorldManager) GetDefaultWorld() world.World {
	return w.defaultWorld
}
func (w *WorldManager) SetDefaultWorld(world world.World) {
	w.defaultWorld = world
}

func NewWorldManager(defaultWorld string) *WorldManager {
	manager := WorldManager{worlds: make(map[string]world.World, 1)}
	readWorlds(&manager)
	world := manager.worlds[defaultWorld]
	if world != nil {
		manager.defaultWorld = world
	}
	return &manager
}

func (w *WorldManager) GetWorld(name string) world.World {
	return w.worlds[name]
}

func (w *WorldManager) UnloadWorld(name string) {
	delete(w.worlds, name)
}

func (w *WorldManager) LoadWorld(filePath string) {
	worldName := getMinWorldName(filePath)
	if worldName == "" {
		return
	}
	world := readWorld(filePath, worldName)
	if world != nil {
		w.worlds[worldName] = world
	}
}
