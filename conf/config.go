package conf

func CreateDefaultConf() ServerConfig {
	return ServerConfig{
		Network{
			Host:        "0.0.0.0",
			Port:        25565,
			Compression: 256,
		},
		Motd{
			Line:    "&bA Golang server",
			Favicon: "",
			Max:     2024,
			Online:  0,
		},
		Game{
			DefaultWorld: "world",
			OnlineMode:   true,
			SendPlayers:  true,
			DebugPackets: true,
		},
	}
}

type ServerConfig struct {
	Network Network
	Motd    Motd
	Game    Game
}

type Network struct {
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	Compression int    `toml:"compression-threshold"`
}

type Game struct {
	DefaultWorld string `toml:"default-world"`
	SendPlayers  bool   `toml:"send-players"`
	OnlineMode   bool   `toml:"online-mode"`
	DebugPackets bool   `toml:"debug-packets"`
}

type Motd struct {
	Line    string `toml:"motd"`
	Favicon string `toml:"favicon"`

	Max    int `toml:"max"`
	Online int `toml:"online"`
}
