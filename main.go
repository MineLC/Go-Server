package main

import (
	"io"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"

	"github.com/golangmc/minecraft-server/impl"
	"github.com/golangmc/minecraft-server/impl/conf"
)

func main() {
	color.NoColor = false

	conf.Config = startConfig()

	server := impl.NewServer()
	server.Load()

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
