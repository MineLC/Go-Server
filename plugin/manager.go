package plugin

import (
	"github.com/minelc/go-server-api/cmd"
	"github.com/minelc/go-server-api/plugin"
)

type listenerData struct {
	plugin plugin.Plugin
	handle func(interface{})
}

type PluginManager struct {
	plugins map[string]plugin.Plugin
	cmd     cmd.CommandManager

	listeners map[int32][]*listenerData
}

func NewPluginManager(cmd cmd.CommandManager) plugin.PluginManager {
	return &PluginManager{
		plugins:   make(map[string]plugin.Plugin, 1),
		cmd:       cmd,
		listeners: make(map[int32][]*listenerData, 1),
	}
}

func (p *PluginManager) GetPlugin(name string) plugin.Plugin {
	return p.plugins[name]
}

func (p *PluginManager) GetCommandManager() cmd.CommandManager {
	return p.cmd
}

func (p *PluginManager) CallEvent(event interface{}, eventType int32) {
	listeners := p.listeners[eventType]
	if listeners == nil {
		return
	}
	for _, listener := range listeners {
		listener.handle(event)
	}
}

func (p *PluginManager) AddListener(listener func(interface{}), eventType int32, plugin plugin.Plugin) {
	eventListeners := p.listeners[eventType]
	data := listenerData{
		plugin: plugin,
		handle: listener,
	}

	if eventListeners == nil {
		listeners := make([]*listenerData, 1)
		listeners[0] = &data
		p.listeners[eventType] = listeners
		return
	}
	newListeners := append(eventListeners, &data)
	p.listeners[eventType] = newListeners
}

func (p *PluginManager) RemoveListener(eventType int32, plugin plugin.Plugin) {
	listeners := p.listeners[eventType]
	if listeners == nil {
		return
	}
	length := len(listeners)

	newSize := 0
	newArray := make([]*listenerData, length)
	position := 0

	for i := 0; i < length; i++ {
		listener := listeners[i]
		if listener.plugin.Name() != plugin.Name() {
			newSize++
			newArray[position] = listener
			position++
			continue
		}
	}

	resized := newArray[:newSize]
	p.listeners[eventType] = resized
}
