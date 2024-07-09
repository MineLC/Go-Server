package main

import (
	"io"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/minelc/go-server/api/data/motd"
	"github.com/minelc/go-server/impl"
	"github.com/minelc/go-server/impl/conf"
	"github.com/minelc/go-server/impl/network"
)

func main() {
	conf := startConfig()
	err := network.StartNet(conf.Network.Port, conf.Network.Host)
	if err != nil {
		return
	}
	stop := make(chan bool)
	impl.Start(stop)
	motd.SetResponse(conf.Motd)

	for {
		if <-stop {
			return
		}
		// TODO: Main thread handler. Use: Syncronize all parts of the project
		// Works in ticks. 1 tick = 50ms
	}
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
