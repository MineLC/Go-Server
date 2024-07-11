package main

import (
	"io"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/minelc/go-server-api/data"
	"github.com/minelc/go-server-api/data/motd"
	"github.com/minelc/go-server/conf"
	"github.com/minelc/go-server/network"
	server "github.com/minelc/go-server/server"
)

func main() {
	conf := startConfig()
	m := conf.Motd

	data.Server = data.ServerData{
		Motd: motd.CreateResponse(motd.Motd{
			VersionName: "1.8.8",
			Line:        m.Line,
			Favicon:     m.Favicon,
			Protocol:    47,
			MaxPlayers:  m.Max,
			Online:      m.Online,
			Sample:      []motd.SamplePlayer{},
		}),
	}

	srv := server.Start()

	err := network.StartNet(conf.Network.Port, conf.Network.Host, srv.GetPackets())
	if err != nil {
		panic(err.Error())
	}

	srv.LoadPlugins()
	server.StartMainLoop(srv)
}

func startConfig() conf.ServerConfig {
	file, err := os.Open("config.toml")

	if file == nil {
		data, err := toml.Marshal(conf.DefaultServerConfig)

		if err != nil {
			panic(err.Error())
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			panic(err.Error())
		}
		return conf.DefaultServerConfig
	}

	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}

	var config conf.ServerConfig

	if err := toml.Unmarshal(b, &config); err != nil {
		panic(err.Error())
	}
	return config
}
