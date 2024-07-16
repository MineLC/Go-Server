package plugin

import (
	"fmt"
	"io/fs"
	"os"
	"plugin"
	"strings"

	api "github.com/minelc/go-server-api"
	api_plugin "github.com/minelc/go-server-api/plugin"
)

var plugins map[string]api_plugin.Plugin

func LoadPlugins(srv api.Server) int {
	files, err := os.ReadDir("plugins")
	if err != nil {
		return 0
	}

	plugins = make(map[string]api_plugin.Plugin, 1)

	for _, file := range files {
		func() {
			srv.Broadcast(file.Name())
			loadPlugin(file, srv)
		}()
	}
	return len(files)
}

func loadPlugin(file fs.DirEntry, srv api.Server) {
	fileName := file.Name()
	if len(fileName) < 4 || !strings.Contains(fileName, ".so") {
		return
	}
	so, err := plugin.Open("plugins/" + fileName)

	if err != nil {
		srv.GetConsole().SendMsgColor("&cError on open the plugin "+fileName, err.Error())
		return
	}
	value, err := so.Lookup("Plugin")
	if err != nil {
		srv.GetConsole().SendMsgColor("&cError on starting the plugin "+fileName, err.Error())
		return
	}
	pl, ok := value.(api_plugin.Plugin)
	if !ok {
		srv.GetConsole().SendMsgColor("&cThe file &6" + fileName + "&c is not a plugin")
		return
	}
	name := pl.Name()
	srv.GetConsole().SendMsgColor("&aStarting plugin... &6" + name)

	defer func() {
		if r := recover(); r != nil {
			srv.GetConsole().SendMsgColor(fmt.Sprint("&cError executing the plugin. Error:", r))
			return
		}
	}()

	pl.Enable()
	plugins[name] = pl
}

func StopPlugins(srv api.Server, complete chan bool) {
	for _, pl := range plugins {
		func() {
			defer func() {
				if r := recover(); r != nil {
					srv.GetConsole().SendMsgColor(fmt.Sprint("&cError stopping the plugin. Error:", r))
					return
				}
			}()
			pl.Disable()
		}()
	}
	complete <- true
}
