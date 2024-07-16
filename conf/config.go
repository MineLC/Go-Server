package conf

var DefaultServerConfig = ServerConfig{
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
	Settings{
		DefaultWorld: "world",
	},
}

type ServerConfig struct {
	Network  Network
	Motd     Motd
	Settings Settings
}

type Network struct {
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	Compression int    `toml:"compression-threshold"`
}

type Settings struct {
	DefaultWorld string `toml:"default-world"`
}

type Motd struct {
	Line    string `toml:"motd"`
	Favicon string `toml:"favicon"`

	Max    int `toml:"max"`
	Online int `toml:"online"`
}
